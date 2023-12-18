package usecases

import (
	"strings"

	"github.com/eggysetiawan/fiber-starterkit/errs"
	"github.com/eggysetiawan/fiber-starterkit/internal/domain"
	"github.com/eggysetiawan/fiber-starterkit/internal/dto"
)

type MachineUseCase interface {
	// GetMachine(id string) (*dto.MachineResponse, *errs.AppError)
	GetMachine(id string) (*dto.MachineResponse, *errs.AppError)
}

type DefaultMachineUseCase struct {
	repository domain.MachineRepository
}

func NewMachineUseCase(repository domain.MachineRepository) DefaultMachineUseCase {
	return DefaultMachineUseCase{repository}
}

func (s *DefaultMachineUseCase) GetMachine(id string) (*dto.MachineResponse, *errs.AppError) {
	identifier := "tid"

	if strings.Contains(id, ".") {
		identifier = "ip_address"
	}

	m, err := s.repository.FindBy(identifier, id)

	if err != nil {
		return nil, err
	}

	response := m.ToDto()

	return response, nil
}
