package config

import (
	"os"
	"strconv"

	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/helper"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/model"
	"github.com/gofiber/fiber/v2"
)

func NewFiber() (*fiber.App, error) {
	prefork, err := strconv.ParseBool(os.Getenv("FIBER_PREFORK"))
	if err != nil {
		return nil, err
	}

	app := fiber.New(fiber.Config{
		Prefork: prefork,
		AppName: os.Getenv("APP_NAME"),
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if e, ok := err.(*helper.AppError); ok {
				return ctx.Status(e.Code).JSON(e)
			}

			return ctx.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
				Code:    fiber.StatusInternalServerError,
				Message: err.Error(),
				Status:  "INTERNAL SERVER ERROR",
			})
		},
	})

	return app, nil
}
