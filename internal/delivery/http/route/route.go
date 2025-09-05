package route

import (
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/delivery/http/middleware"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/module/auth"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/module/user"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App *fiber.App
	*auth.AuthController
	*user.UserController
	*middleware.Middleware
}

func (c *RouteConfig) Setup() {
	api := c.App.Group("/api")
	v1 := api.Group("/v1")

	c.AuthHandler(v1)
	c.UserHandler(v1)
}

func (c *RouteConfig) AuthHandler(router fiber.Router) {
	r := router.Group("/auth")
	r.Post("/login", c.AuthController.Login)
	r.Post("/register", c.AuthController.Register)
	r.Post("/refresh", c.AuthController.Refresh)
}

func (c *RouteConfig) UserHandler(router fiber.Router) {
	users := router.Group("/users", c.Middleware.AuthMiddleware)
	users.Get("/me", c.UserController.Me)
	users.Get("/:id", c.Middleware.RoleMiddleware("admin"), c.UserController.User)
}
