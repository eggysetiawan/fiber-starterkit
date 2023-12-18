package helper

import (
	"github.com/eggysetiawan/fiber-starterkit/internal/domain"
	"github.com/gofiber/fiber/v2"
)

func ResponseApi(c *fiber.Ctx, response domain.WebResponse) {
	c.Status(fiber.StatusInternalServerError).JSON(&response)
}

func ResponseApiError(c *fiber.Ctx, response domain.ResponseError) {
	c.Status(fiber.StatusInternalServerError).JSON(&response)
}
