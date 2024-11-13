package handlers

import (
	"context"
	"time"

	"github.com/eggysetiawan/fiber-starterkit/config"
	"github.com/eggysetiawan/fiber-starterkit/internal/domain"
	"github.com/eggysetiawan/fiber-starterkit/internal/dto"
	"github.com/eggysetiawan/fiber-starterkit/internal/usecases"
	"github.com/eggysetiawan/fiber-starterkit/logger"
	"github.com/gofiber/fiber/v2"
)

type BigDataHandler struct {
	uc *usecases.IBigDataUseCase
}

func NewBigDataHandler(uc usecases.IBigDataUseCase) *BigDataHandler {
	return &BigDataHandler{&uc}
}

func (h *BigDataHandler) GetPostalCode(c *fiber.Ctx) error {
	request := dto.PostalCodeRequest{}

	if err := c.BodyParser(&request); err != nil {
		logger.Error(err.Error())
		return domain.NewBadRequestResponse(c)
	}

	d := time.Second * config.AppConfig.GetDuration("BRIGATE_TIMEOUT_SECOND")
	ctx, cancel := context.WithTimeout(c.Context(), d)
	defer cancel()

	data, err := h.uc.GetPostalCodes(ctx, request)
	if err != nil {
		return domain.NewUnexpectedErrorResponse(c, err.Message)
	}

	select {
	case <-ctx.Done():
		if ctx.Err() == context.DeadlineExceeded {
			return domain.NewTimeoutResponse(c)
		}
	default:
	}

	return domain.NewSuccessfulResponse(c, data)

}
