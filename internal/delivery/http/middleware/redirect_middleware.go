package middleware

import (
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/helper"
	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) RedirectMiddleware(ctx *fiber.Ctx) error {
	redirectUri := ctx.Query("redirect_uri")

	isRedirectValidUri := helper.ValidRedirectUrl(redirectUri)

	if !isRedirectValidUri {
		return helper.BadRequest("Unknown service redirect")
	}

	return ctx.Next()
}
