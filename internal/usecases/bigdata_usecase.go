package usecases

import (
	"context"
	"sync"

	"github.com/eggysetiawan/fiber-starterkit/errs"
	"github.com/eggysetiawan/fiber-starterkit/internal/domain"
	"github.com/eggysetiawan/fiber-starterkit/internal/dto"
)

type IBigDataUseCase struct {
	repo domain.IBigDataRepository
}

func NewDefaultBigDataUseCase(repo domain.IBigDataRepository) *IBigDataUseCase {
	return &IBigDataUseCase{repo}
}

func (uc *IBigDataUseCase) GetPostalCodes(c context.Context, request dto.PostalCodeRequest) (*dto.WithPaginationResponse[dto.PostalCodeResponse], *errs.AppError) {
	uniqueMap := make(map[int]struct{}, 0)

	postalCodes := make([]int, 0)

	for _, pc := range request.PostalCode {
		if _, exists := uniqueMap[pc]; !exists {
			postalCodes = append(postalCodes, pc)
			uniqueMap[pc] = struct{}{}
		}
	}

	ch := make(chan domain.PostalCode, len(postalCodes))
	errCh := make(chan *errs.AppError, len(postalCodes))
	wg := &sync.WaitGroup{}

	wg.Add(len(postalCodes))
	for _, pc := range postalCodes {
		go uc.repo.GetPostalCode(c, pc, ch, wg, errCh)
	}
	wg.Wait()

	close(ch)
	close(errCh)

	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	data := make([]dto.PostalCodeResponse, 0)

	for postalCode := range ch {
		response := postalCode.ToDto()
		data = append(data, response)
	}

	resp := dto.WithPaginationResponse[dto.PostalCodeResponse]{
		Data: data,
	}

	return &resp, nil
}
