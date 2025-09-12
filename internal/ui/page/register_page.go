package page

import (
	"net/url"
	"os"

	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/constant"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/helper"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/model"
	"github.com/gofiber/fiber/v2"
)

func (p *Pages) RegisterHandleForm(ctx *fiber.Ctx) error {
	redirectUri := ctx.Query("redirect_uri")

	request := model.AuthRegisterRequest{}

	params := url.Values{}
	u, _ := url.Parse(p.Config.Prefix)
	params.Add("redirect_uri", redirectUri)

	if redirectUri == "" {
		return ctx.Redirect(os.Getenv("ORDER_SERVICE_URL"))
	}

	if err := ctx.BodyParser(&request); err != nil {
		params.Add("error", err.Error())
		params.Add("email", request.Email)
		params.Add("name", request.Name)
		u.RawQuery = params.Encode()
		return ctx.Redirect(u.String())
	}

	response, err := p.Auth.AuthUseCase.Register(ctx.UserContext(), request)
	if err != nil {
		params.Add("error", err.Error())
		params.Add("email", request.Email)
		params.Add("name", request.Name)
		u.RawQuery = params.Encode()
		return ctx.Redirect(u.String())
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     constant.REFRESH_TOKEN,
		Value:    response.RefreshToken,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
		MaxAge:   constant.MAX_AGE_REFRESH_TOKEN,
		Path:     "/",
	})

	ctx.Cookie(&fiber.Cookie{
		Name:     constant.ACCESS_TOKEN,
		Value:    response.AccessToken,
		HTTPOnly: false,
		Secure:   true,
		SameSite: "Lax",
		MaxAge:   constant.MAX_AGE_ACCESS_TOKEN,
		Path:     "/",
	})

	return ctx.Redirect(redirectUri)
}

func (p *Pages) PageRegister(ctx *fiber.Ctx) error {
	redirectUri := ctx.Query("redirect_uri")
	email := ctx.Query("email")
	errorMessage := ctx.Query("error")
	name := ctx.Query("name")

	actionPostUrl := helper.BuildURL(
		p.Config.Prefix+"/register",
		map[string]string{
			"redirect_uri": redirectUri,
		},
	)

	signInUrl := helper.BuildURL(
		p.Config.Prefix+"/",
		map[string]string{
			"redirect_uri": redirectUri,
		},
	)

	return ctx.Render("page/register", fiber.Map{
		"Title":         "Shophub Back!",
		"Message":       errorMessage,
		"Name":          name,
		"Email":         email,
		"ActionPostUrl": actionPostUrl,
		"SignInUrl":     signInUrl,
	})
}
