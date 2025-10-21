package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"polimane/backend/env"
)

const (
	cookieAccessToken  = "pat"
	cookieRefreshToken = "prt"
)

var (
	cookieMaxAge = int((time.Hour * 24 * 60).Seconds())
)

type authCookies struct {
	AccessToken  string `cookie:"pat"`
	RefreshToken string `cookie:"prt"`
}

func getCookies(ctx *fiber.Ctx) (*authCookies, error) {
	var cookies authCookies
	if err := ctx.CookieParser(&cookies); err != nil {
		return nil, err
	}
	return &cookies, nil
}

func setCookies(ctx *fiber.Ctx, environment *env.Environment, cookies *authCookies) {
	domain := "." + environment.AppDomain

	accessToken := createCookie(domain, cookieAccessToken, cookies.AccessToken)
	accessToken.MaxAge = cookieMaxAge
	ctx.Cookie(accessToken)

	refreshToken := createCookie(domain, cookieRefreshToken, cookies.RefreshToken)
	refreshToken.MaxAge = cookieMaxAge
	ctx.Cookie(refreshToken)
}

func removeCookies(ctx *fiber.Ctx, environment *env.Environment) {
	domain := "." + environment.AppDomain

	accessToken := createCookie(domain, cookieAccessToken, "-")
	accessToken.Expires = time.Unix(0, 0)
	ctx.Cookie(accessToken)

	refreshToken := createCookie(domain, cookieRefreshToken, "-")
	refreshToken.Expires = time.Unix(0, 0)
	ctx.Cookie(refreshToken)
}

func createCookie(domain, name, value string) *fiber.Cookie {
	return &fiber.Cookie{
		Name:     name,
		Value:    value,
		Domain:   domain,
		Secure:   true,
		HTTPOnly: true,
		SameSite: "None",
	}
}
