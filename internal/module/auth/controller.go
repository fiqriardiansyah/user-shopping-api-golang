package auth

import (
	"strings"

	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/delivery/http/middleware"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/helper"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/model"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/module/auth/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	*usecase.AuthUseCase
	*validator.Validate
	*helper.Config
}

func NewAuthController(authUseCase *usecase.AuthUseCase, validator *validator.Validate, config *helper.Config) *AuthController {
	return &AuthController{
		AuthUseCase: authUseCase,
		Validate:    validator,
		Config:      config,
	}
}

func (c *AuthController) RegisterRoutes(router fiber.Router, mw *middleware.Middleware) {
	r := router.Group("/auth")
	r.Post("/login", c.Login)
	r.Post("/register", c.Register)
	r.Post("/refresh", c.Refresh)
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	request := model.AuthRegisterRequest{}

	if err := ctx.BodyParser(&request); err != nil {
		return helper.BadRequest(err.Error())
	}

	if err := c.Validate.Struct(&request); err != nil {
		return helper.BadRequest(err.Error())
	}

	response, err := c.AuthUseCase.Register(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return helper.Success(ctx, response, 201)
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	request := model.AuthLoginRequest{}

	if err := ctx.BodyParser(&request); err != nil {
		return helper.BadRequest(err.Error())
	}

	if err := c.Validate.Struct(&request); err != nil {
		return helper.BadRequest(err.Error())
	}

	response, err := c.AuthUseCase.Login(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return helper.Success(ctx, response, 200)
}

func (c *AuthController) Refresh(ctx *fiber.Ctx) error {
	authorization := ctx.Get("Authorization")
	if authorization == "" || !strings.HasPrefix(authorization, "Bearer ") {
		return helper.Unauthorized("Authorization token not valid!")
	}
	token := strings.TrimPrefix(authorization, "Bearer ")

	response, err := c.AuthUseCase.Refresh(token)
	if err != nil {
		return err
	}

	return helper.Success(ctx, response, 200)
}
