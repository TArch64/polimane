package auth

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kittipat1413/go-common/framework/cache"
	"github.com/kittipat1413/go-common/framework/cache/localcache"
	"github.com/workos/workos-go/v4/pkg/usermanagement"
	"go.uber.org/fx"
	"gorm.io/gorm"

	"polimane/backend/api/base"
	"polimane/backend/env"
	"polimane/backend/model"
	repositoryusers "polimane/backend/repository/users"
	"polimane/backend/services/workos"
	"polimane/backend/signal"
)

var (
	unauthorizedErr   = base.NewReasonedError(fiber.StatusUnauthorized, "Unauthorized")
	missingTokenErr   = errors.New("missing access or refresh token")
	sessionExpiredErr = errors.New("session expired")
)

type MiddlewareOptions struct {
	fx.In
	Signals      *signal.Container
	Env          *env.Environment
	WorkosClient workos.Client
	Users        repositoryusers.Client
}

type Middleware struct {
	userCache       cache.Cache[*model.User]
	workosUserCache cache.Cache[*usermanagement.User]
	workosClient    workos.Client
	env             *env.Environment
	users           repositoryusers.Client
}

func MiddlewareProvider(options MiddlewareOptions) *Middleware {
	cacheOptions := []localcache.Option{
		localcache.WithDefaultExpiration(10 * time.Minute),
		localcache.WithCleanupInterval(5 * time.Minute),
	}

	middleware := &Middleware{
		userCache:       localcache.New[*model.User](cacheOptions...),
		workosUserCache: localcache.New[*usermanagement.User](cacheOptions...),
		env:             options.Env,
		workosClient:    options.WorkosClient,
		users:           options.Users,
	}

	options.Signals.InvalidateUserCache.AddListener(middleware.invalidateUserCache)
	options.Signals.InvalidateWorkosUserCache.AddListener(middleware.invalidateWorkosUserCache)
	options.Signals.InvalidateAuthCache.AddListener(middleware.invalidateAuthCache)
	return middleware
}

func (m *Middleware) invalidateAuthCache(ctx context.Context, sessionID string) {
	workosUser, _ := m.workosUserCache.Get(ctx, sessionID, nil)

	if workosUser != nil {
		m.invalidateWorkosUserCache(ctx, workosUser.ID)
		m.invalidateUserCache(ctx, model.MustStringToID(workosUser.ExternalID))
	}
}

func (m *Middleware) invalidateUserCache(ctx context.Context, userID model.ID) {
	_ = m.userCache.Invalidate(ctx, userID.String())
}

func (m *Middleware) invalidateWorkosUserCache(ctx context.Context, userID string) {
	_ = m.workosUserCache.Invalidate(ctx, userID)
}

func (m *Middleware) Handler(ctx *fiber.Ctx) error {
	accessToken := ctx.Get("Authorization")
	refreshToken := ctx.Get("X-Refresh-Token")
	if accessToken == "" || refreshToken == "" {
		return m.newUnauthorizedErr(missingTokenErr)
	}

	accessTokenClaims, err := m.workosClient.AuthenticateWithAccessToken(ctx.Context(), accessToken)
	if errors.Is(err, workos.AccessTokenExpired) {
		accessTokenClaims, err = m.refreshToken(ctx, refreshToken)
	}
	if err != nil {
		return err
	}

	workosUser, err := m.getWorkosUser(ctx.Context(), accessTokenClaims)
	if err != nil {
		return err
	}

	userID, err := model.StringToID(workosUser.ExternalID)
	if err != nil {
		return err
	}

	user, err := m.getUser(ctx.Context(), userID)
	if err != nil {
		return err
	}

	SetSession(ctx, &UserSession{
		User:       user,
		WorkosUser: workosUser,
		ID:         accessTokenClaims.SessionID,
	})

	return ctx.Next()
}

func (m *Middleware) refreshToken(ctx *fiber.Ctx, token string) (*workos.AccessTokenClaims, error) {
	res, err := m.workosClient.UserManagement().AuthenticateWithRefreshToken(ctx.Context(), usermanagement.AuthenticateWithRefreshTokenOpts{
		ClientID:     m.env.WorkOS.ClientID,
		RefreshToken: token,
		UserAgent:    ctx.Get("User-Agent"),
	})

	if err != nil {
		if strings.Contains(err.Error(), "Session ended") {
			return nil, m.newUnauthorizedErr(sessionExpiredErr)
		}
		return nil, err
	}

	ctx.Set("X-New-Refresh-Token", res.RefreshToken)
	ctx.Set("X-New-Access-Token", res.AccessToken)
	return m.workosClient.AuthenticateWithAccessToken(ctx.Context(), res.AccessToken)
}

func (m *Middleware) getWorkosUser(ctx context.Context, accessTokenClaims *workos.AccessTokenClaims) (*usermanagement.User, error) {
	return m.workosUserCache.Get(ctx, accessTokenClaims.UserID, func() (*usermanagement.User, *time.Duration, error) {
		user, err := m.workosClient.UserManagement().GetUser(ctx, usermanagement.GetUserOpts{
			User: accessTokenClaims.UserID,
		})

		if err != nil {
			return nil, nil, err
		}

		return &user, nil, nil
	})
}

func (m *Middleware) getUser(ctx context.Context, id model.ID) (*model.User, error) {
	return m.userCache.Get(ctx, id.String(), func() (*model.User, *time.Duration, error) {
		user, err := m.users.ByID(ctx, id)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, m.newUnauthorizedErr(err, base.CustomErrorData{
				"userId": id.String(),
			})
		}
		if err != nil {
			return nil, nil, err
		}

		return user, nil, nil
	})
}

func (m *Middleware) newUnauthorizedErr(err error, extra ...base.CustomErrorData) error {
	if env.IsDev {
		extra = append(extra, base.CustomErrorData{"internalError": err.Error()})
		return unauthorizedErr.AddCustomData(extra...)
	}

	return unauthorizedErr
}
