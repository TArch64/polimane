package auth

import (
	"context"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/guregu/dynamo/v2"
	"github.com/kittipat1413/go-common/framework/cache"
	"github.com/kittipat1413/go-common/framework/cache/localcache"

	"polimane/backend/api/base"
	"polimane/backend/env"
	"polimane/backend/model"
	repositoryusers "polimane/backend/repository/users"
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

func (m *middleware) invalidateCache(ctx context.Context, userID model.ID) {
	_ = m.cache.Invalidate(ctx, userID.String())
}

func (m *middleware) Handler(ctx *fiber.Ctx) error {
	token, err := m.getToken(ctx)
	if err != nil {
		return unauthorizedErr
	}

	claims, err := m.parseToken(token)
	if err != nil {
		return err
	}

	user, err := m.getUser(ctx.Context(), claims)
	if err != nil {
		return err
	}

	setSessionUser(ctx, user)
	return ctx.Next()
}

func (m *middleware) getToken(ctx *fiber.Ctx) (string, error) {
	tokens := ctx.GetReqHeaders()["Authorization"]
	if len(tokens) == 0 || len(tokens[0]) == 0 {
		return "", unauthorizedErr
	}
	return tokens[0], nil
}

func (m *middleware) parseToken(token string) (*tokenClaims, error) {
	claims := &tokenClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(env.Env().SecretKey), nil
	})

	return claims, err
}

func (m *middleware) getUser(ctx context.Context, claims *tokenClaims) (*model.User, error) {
	return m.cache.Get(ctx, claims.UserID.String(), func() (*model.User, *time.Duration, error) {
		user, err := repositoryusers.ByID(ctx, claims.UserID)

		if errors.Is(err, dynamo.ErrNotFound) {
			return nil, nil, unauthorizedErr
		}
		if err != nil {
			return nil, nil, err
		}

		return user, nil, nil
	})
}
