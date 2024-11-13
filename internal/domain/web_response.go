package domain

import "github.com/gofiber/fiber/v2"

type WebResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

func NewSuccessfulResponse(c *fiber.Ctx, data interface{}) error {
	c.SendStatus(fiber.StatusOK)
	return c.JSON(WebResponse{
		Message: "Sukses",
		Code:    fiber.StatusOK,
		Data:    data,
	})
}

func NewBadRequestResponse(c *fiber.Ctx) error {
	c.SendStatus(fiber.StatusBadRequest)
	return c.JSON(WebResponse{
		Message: "Bad Request",
		Code:    fiber.StatusBadRequest,
		Data:    []string{},
	})
}

func NewUnexpectedErrorResponse(c *fiber.Ctx, m string) error {
	c.SendStatus(fiber.StatusInternalServerError)

	return c.JSON(WebResponse{
		Message: m,
		Code:    fiber.StatusInternalServerError,
		Data:    []string{},
	})
}

func NewUnauthorizedResponse(c *fiber.Ctx) error {
	c.SendStatus(fiber.StatusUnauthorized)

	return c.JSON(WebResponse{
		Message: "Unauthorized",
		Code:    fiber.StatusUnauthorized,
		Data:    []string{},
	})
}

func NewTimeoutResponse(c *fiber.Ctx) error {
	c.SendStatus(fiber.StatusRequestTimeout)
	return c.JSON(WebResponse{
		Message: "Request too long. got timeout!",
		Code:    fiber.StatusRequestTimeout,
		Data:    []string{},
	})
}

type ResponseError struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Error   interface{} `json:"error"`
}
