package handlers

import (
	"github.com/eggysetiawan/fiber-starterkit/internal/domain"
	"github.com/eggysetiawan/fiber-starterkit/internal/dto"
	"github.com/eggysetiawan/fiber-starterkit/internal/usecases"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type MachineHandler struct {
	uc usecases.DefaultMachineUseCase
}

func NewMachineHandler(uc usecases.DefaultMachineUseCase) *MachineHandler {
	return &MachineHandler{
		uc: uc,
	}
}

func (mh *MachineHandler) ShowMachine(c *fiber.Ctx) error {
	request := new(dto.MachineRequest)

	if err := c.BodyParser(request); err != nil {
		return domain.NewBadRequestResponse(c)
	}

	data, err := mh.uc.GetMachine(request.Identifier)

	if err != nil {
		log.Error(err.Message)
		return c.Status(err.Code).JSON(domain.WebResponse{
			Message: err.Message,
			Code:    err.Code,
			Data:    []string{},
		})
	}

	return domain.NewSuccessfulResponse(c, data)
}
