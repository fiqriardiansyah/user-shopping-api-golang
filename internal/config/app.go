package config

import (
	grpcServers "github.com/fiqriardiansyah/user-shopping-api-golang/internal/delivery/grpc"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/delivery/http/middleware"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/delivery/http/route"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/helper"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/module/auth"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/module/user"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/ui/page"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type AppConfig struct {
	App        *fiber.App
	Db         *gorm.DB
	Validate   *validator.Validate
	Config     *helper.Config
	GrpcServer *grpc.Server
}

func NewApp(config *AppConfig) {
	authHandlerSet := auth.InitializeAuthHandler(config.Db, config.Validate, config.Config)
	userHandlerSet := user.InitializeUserHandler(config.Db, config.Validate, config.Config)

	page := page.NewPages(authHandlerSet.Controller, userHandlerSet.Controller, config.Config)

	// grpc handler
	userGrpc := grpcServers.NewUserServer(userHandlerSet.UseCase)
	userGrpc.Run(config.GrpcServer)

	r := route.RouteConfig{
		App:        config.App,
		Auth:       authHandlerSet.Controller,
		User:       userHandlerSet.Controller,
		Middleware: middleware.NewMiddleware(config.Db),
		Page:       page,
	}

	r.Setup()
}
