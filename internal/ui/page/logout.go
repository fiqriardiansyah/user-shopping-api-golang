package page

import (
	"os"
	"time"

	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/constant"
	"github.com/gofiber/fiber/v2"
)

func (c *Pages) Logout(ctx *fiber.Ctx) error {
	redirectUri := ctx.Query("redirect_uri")

	if redirectUri == "" {
		return ctx.Redirect(os.Getenv("ORDER_SERVICE_URL"))
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     constant.ACCESS_TOKEN,
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})

	ctx.Cookie(&fiber.Cookie{
		Name:     constant.REFRESH_TOKEN,
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})

	return ctx.Redirect(redirectUri)
}
