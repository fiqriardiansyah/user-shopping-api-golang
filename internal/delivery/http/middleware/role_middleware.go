package middleware

import (
	"slices"

	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/entity"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/helper"
	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) RoleMiddleware(roles ...string) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		user := ctx.Locals("user").(*entity.User)

		for _, role := range user.Roles {
			if slices.Contains(roles, role.Name) {
				return ctx.Next()
			}
		}

		return helper.Forbidden("Forbidden resource")
	}
}
