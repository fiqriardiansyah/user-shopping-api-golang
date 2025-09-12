package config

import (
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/delivery/http/middleware"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/delivery/http/route"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/helper"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/module/auth"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/module/user"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/ui/page"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AppConfig struct {
	App      *fiber.App
	Db       *gorm.DB
	Validate *validator.Validate
	Config   *helper.Config
}

func NewApp(config *AppConfig) {
	authController := auth.InitializeAuthHandler(config.Db, config.Validate, config.Config)
	userController := user.InitializeUserHandler(config.Db, config.Validate, config.Config)

	page := page.NewPages(authController, userController, config.Config)

	r := route.RouteConfig{
		App:        config.App,
		Auth:       authController,
		User:       userController,
		Middleware: middleware.NewMiddleware(config.Db),
		Page:       page,
	}

	r.Setup()
}
