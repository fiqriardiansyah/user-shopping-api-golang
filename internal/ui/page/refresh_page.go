package page

import (
	"os"

	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/constant"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/helper"
	"github.com/gofiber/fiber/v2"
)

func (p *Pages) RefreshHandler(ctx *fiber.Ctx) error {
	redirectUri := ctx.Query("redirect_uri")
	if redirectUri == "" {
		return ctx.Redirect(os.Getenv("ORDER_SERVICE_URL"))
	}

	refreshToken := ctx.Cookies("refresh_token")

	_, err := helper.ValidateToken(refreshToken, os.Getenv("JWT_REFRESH_TOKEN_SECRET"))
	if err != nil {
		return ctx.Redirect("/logout?redirect_uri=" + redirectUri)
	}

	response, err := p.Auth.AuthUseCase.Refresh(refreshToken)
	if err != nil {
		return ctx.Redirect("/?error=" + err.Error())
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     constant.REFRESH_TOKEN,
		Value:    response.RefreshToken,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
		MaxAge:   constant.MAX_AGE_REFRESH_TOKEN,
		Path:     "/",
	})

	ctx.Cookie(&fiber.Cookie{
		Name:     constant.ACCESS_TOKEN,
		Value:    response.AccessToken,
		HTTPOnly: false,
		Secure:   true,
		SameSite: "Lax",
		MaxAge:   constant.MAX_AGE_ACCESS_TOKEN,
		Path:     "/",
	})

	return ctx.Redirect(redirectUri)
}
