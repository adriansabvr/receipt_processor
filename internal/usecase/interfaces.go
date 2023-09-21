package usecase

import (
	"context"
	"github.com/adriansabvr/receipt_processor/internal/entity"
)

type (
	// Receipt -.
	Receipt interface {
		Process(context.Context, entity.Receipt) (uint64, error)
		GetPoints(context.Context, uint64) (int, error)
	}

	// ReceiptRepo -.
	ReceiptRepo interface {
		InsertReceipt(context.Context, entity.Receipt) (uint64, error)
		GetReceipt(context.Context, uint64) (entity.Receipt, error)
	}
)
