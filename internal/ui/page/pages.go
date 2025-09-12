package page

import (
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/delivery/http/middleware"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/helper"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/module/auth"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/module/user"
	"github.com/gofiber/fiber/v2"
)

type Pages struct {
	Auth   *auth.AuthController
	User   *user.UserController
	Config *helper.Config
}

func NewPages(auth *auth.AuthController, user *user.UserController, config *helper.Config) *Pages {
	return &Pages{
		Auth:   auth,
		User:   user,
		Config: config,
	}
}

func (p *Pages) RegisterPages(router fiber.Router, mw *middleware.Middleware) {
	router.Get("/", mw.RedirectMiddleware, p.PageLogin)
	router.Post("/", p.LoginHandleForm)
	router.Get("/register", mw.RedirectMiddleware, p.PageRegister)
	router.Post("/register", p.RegisterHandleForm)
	router.Get("/refresh", p.RefreshHandler)
	router.Get("/logout", p.Logout)
}
