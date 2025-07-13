package auth

import (
	"context"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kittipat1413/go-common/framework/cache"
	"github.com/kittipat1413/go-common/framework/cache/localcache"
	"github.com/workos/workos-go/v4/pkg/usermanagement"
	"gorm.io/gorm"

	"polimane/backend/api/base"
	"polimane/backend/env"
	"polimane/backend/model"
	"polimane/backend/model/modelbase"
	repositoryusers "polimane/backend/repository/users"
	"polimane/backend/services/workos"
	"polimane/backend/signal"
)

var unauthorizedErr = base.NewReasonedError(fiber.StatusUnauthorized, "Unauthorized")

type middleware struct {
	cache cache.Cache[*model.User]
}

func NewMiddleware() fiber.Handler {
	m := &middleware{
		cache: localcache.New[*model.User](
			localcache.WithDefaultExpiration(10*time.Minute),
			localcache.WithCleanupInterval(5*time.Minute),
		),
	}

	signal.InvalidateAuthCache.AddListener(m.invalidateCache)
	return m.Handler
}

func (m *middleware) invalidateCache(ctx context.Context, userID modelbase.ID) {
	_ = m.cache.Invalidate(ctx, userID.String())
}

func (m *middleware) Handler(ctx *fiber.Ctx) error {
	accessToken := ctx.Get("Authorization")
	refreshToken := ctx.Get("X-Refresh-Token")
	if accessToken == "" || refreshToken == "" {
		return unauthorizedErr
	}

	workosUser, err := workos.AuthenticateWithAccessToken(ctx.Context(), accessToken)
	if errors.Is(err, workos.AccessTokenExpired) {
		workosUser, err = m.refreshToken(ctx, refreshToken)
	}
	if err != nil {
		return err
	}

	userID, err := modelbase.StringToID(workosUser.ExternalID)
	if err != nil {
		return err
	}

	user, err := m.getUser(ctx.Context(), userID)
	if err != nil {
		return err
	}

	setSession(ctx, &UserSession{
		User:       user,
		WorkosUser: workosUser,
	})

	return ctx.Next()
}

func (m *middleware) refreshToken(ctx *fiber.Ctx, token string) (*usermanagement.User, error) {
	res, err := workos.UserManagement.AuthenticateWithRefreshToken(ctx.Context(), usermanagement.AuthenticateWithRefreshTokenOpts{
		ClientID:     env.Instance.WorkOS.ClientID,
		RefreshToken: token,
		UserAgent:    ctx.Get("User-Agent"),
	})

	if err != nil {
		return nil, err
	}

	ctx.Set("X-New-Refresh-Token", res.RefreshToken)
	ctx.Set("X-New-Access-Token", res.AccessToken)
	return workos.AuthenticateWithAccessToken(ctx.Context(), res.AccessToken)
}

func (m *middleware) getUser(ctx context.Context, id modelbase.ID) (*model.User, error) {
	return m.cache.Get(ctx, id.String(), func() (*model.User, *time.Duration, error) {
		user, err := repositoryusers.ByID(ctx, id)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, unauthorizedErr
		}
		if err != nil {
			return nil, nil, err
		}

		return user, nil, nil
	})
}
