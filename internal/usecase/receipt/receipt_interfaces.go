package receipt

import (
	"context"
	"github.com/adriansabvr/receipt_processor/internal/entity"
)

//go:generate mockgen -source=receipt_interfaces.go -destination=./mocks_test.go -package=receipt_test

type (
	// UseCaseContract -.
	UseCaseContract interface {
		Process(context.Context, entity.Receipt) uint64
		GetPoints(context.Context, uint64) (int, error)
	}

	// RepoContract -.
	RepoContract interface {
		InsertReceipt(context.Context, entity.Receipt) uint64
		GetReceipt(context.Context, uint64) (entity.Receipt, error)
	}
)
