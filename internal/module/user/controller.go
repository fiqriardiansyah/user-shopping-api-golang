package user

import (
	"strings"

	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/delivery/http/middleware"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/entity"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/helper"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/module/user/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserController struct {
	*usecase.UserUseCase
	*validator.Validate
	*helper.Config
}

func NewUserController(useCase *usecase.UserUseCase, validator *validator.Validate, config *helper.Config) *UserController {
	return &UserController{
		UserUseCase: useCase,
		Validate:    validator,
		Config:      config,
	}
}

func (c *UserController) RegisterRoutes(router fiber.Router, mw *middleware.Middleware) {
	users := router.Group("/users", mw.AuthMiddleware)
	users.Get("/me", c.Me)
	users.Get("/:id", mw.RoleMiddleware("admin"), c.User)
}

func (c *UserController) Me(ctx *fiber.Ctx) error {
	userLocal := ctx.Locals("user").(*entity.User)
	user, err := c.UserUseCase.Me(userLocal.ID)
	if err != nil {
		return err
	}
	return helper.Success(ctx, &user, 200)
}

func (c *UserController) User(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	if strings.Trim(userId, " ") == "" {
		return helper.BadRequest("Parameter user id required")
	}
	userIdParse, err := uuid.Parse(userId)
	if err != nil {
		return err
	}

	user, err := c.UserUseCase.User(userIdParse)
	return helper.Success(ctx, user, 200)
}
