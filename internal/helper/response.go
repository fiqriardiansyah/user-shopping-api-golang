package helper

import (
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/model"
	"github.com/gofiber/fiber/v2"
)

func Success(c *fiber.Ctx, data any, code int) error {
	if code == 0 {
		code = fiber.StatusOK
	}
	return c.Status(code).JSON(model.WebResponse{
		Data:   data,
		Status: "SUCCESS",
		Code:   code,
	})
}
