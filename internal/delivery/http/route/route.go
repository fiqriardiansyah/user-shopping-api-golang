package route

import (
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/delivery/http/middleware"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/module/auth"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/module/user"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/ui/page"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App        *fiber.App
	Auth       *auth.AuthController
	User       *user.UserController
	Middleware *middleware.Middleware
	Page       *page.Pages
}

func (rc *RouteConfig) Setup() {
	// UI
	rc.Page.RegisterPages(rc.App, rc.Middleware)

	// API
	api := rc.App.Group("/api")
	v1 := api.Group("/v1")
	rc.Auth.RegisterRoutes(v1, rc.Middleware)
	rc.User.RegisterRoutes(v1, rc.Middleware)
}
