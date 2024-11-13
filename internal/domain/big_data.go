package domain

import (
	"context"
	"sync"

	"github.com/eggysetiawan/fiber-starterkit/errs"
)

type IBigDataRepository interface {
	GetPostalCode(context.Context, int, chan PostalCode, *sync.WaitGroup, chan *errs.AppError)
}
