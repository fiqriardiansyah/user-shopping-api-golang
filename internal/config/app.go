package config

import (
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/delivery/http/middleware"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/delivery/http/route"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/module/auth"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/module/user"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AppConfig struct {
	App      *fiber.App
	Db       *gorm.DB
	Validate *validator.Validate
}

func NewApp(config *AppConfig) {
	authController := auth.InitializeAuthHandler(config.Db, config.Validate)
	userController := user.InitializeUserHandler(config.Db, config.Validate)

	r := route.RouteConfig{
		App:            config.App,
		AuthController: authController,
		UserController: userController,
		Middleware:     middleware.NewMiddleware(config.Db),
	}

	r.Setup()
}
