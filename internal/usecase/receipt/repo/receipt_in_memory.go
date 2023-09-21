package repo

import (
	"context"

	"github.com/adriansabvr/receipt_processor/internal/entity"
	"github.com/palantir/stacktrace"
)

var errReceiptNotFound = stacktrace.NewError("receipt not found")

// ReceiptRepo -.
type ReceiptRepo struct {
	mp map[uint64]entity.Receipt
}

// New -.
func New() *ReceiptRepo {
	return &ReceiptRepo{mp: make(map[uint64]entity.Receipt)}
}

var id uint64

// InsertReceipt -.
func (r *ReceiptRepo) InsertReceipt(_ context.Context, receipt entity.Receipt) uint64 {
	id++
	r.mp[id] = receipt

	return id
}

// GetReceipt -.
func (r *ReceiptRepo) GetReceipt(_ context.Context, receiptID uint64) (entity.Receipt, error) {
	receipt, ok := r.mp[receiptID]
	if !ok {
		return entity.Receipt{}, errReceiptNotFound
	}

	return receipt, nil
}
