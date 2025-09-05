package middleware

import (
	"errors"
	"os"
	"strings"

	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/entity"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/helper"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (m *Middleware) AuthMiddleware(ctx *fiber.Ctx) error {

	authorization := ctx.Get("Authorization")
	if authorization == "" || !strings.HasPrefix(authorization, "Bearer ") {
		return helper.Unauthorized("Authorization token not valid or not found")
	}
	token := strings.TrimPrefix(authorization, "Bearer ")

	claims, err := helper.ValidateToken(token, os.Getenv("JWT_SECRET"))
	if err != nil {
		return helper.BadRequest(err.Error())
	}

	user := entity.User{}
	//m.db.Where("id = ?", claims.UserId).Preload("Roles").Preload("Products").Preload("Orders").Preload("Reviews").First(&user)
	result := m.db.Where("id = ?", claims.UserId).Preload(clause.Associations).First(&user)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return helper.NotFound("User not found")
	}

	ctx.Locals("user", &user)

	return ctx.Next()
}
