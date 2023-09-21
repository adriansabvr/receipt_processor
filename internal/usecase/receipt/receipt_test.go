package receipt_test

import (
	"context"
	"testing"
	"time"

	"github.com/adriansabvr/receipt_processor/internal/entity"
	"github.com/adriansabvr/receipt_processor/internal/usecase/receipt"
	"github.com/golang/mock/gomock"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

type test struct {
	id      uint64
	name    string
	mock    func()
	receipt entity.Receipt
	res     interface{}
	err     error
}

func receiptTestService(t *testing.T) (*receipt.UseCase, *MockRepoContract) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := NewMockRepoContract(mockCtl)
	useCase := receipt.New(repo)

	return useCase, repo
}

// getTestReceipt1 returns a test receipt using the examples in the README.md.
func getTestReceipt1() entity.Receipt {
	return entity.Receipt{
		Retailer:     "Target",
		PurchaseDate: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		PurchaseTime: time.Date(0, 1, 1, 13, 1, 0, 0, time.UTC),
		Total:        decimal.NewFromFloat(35.35),
		Items: []entity.Item{
			{
				ShortDescription: "Mountain Dew 12PK",
				Price:            decimal.NewFromFloat(6.49),
			},
			{
				ShortDescription: "Emils Cheese Pizza",
				Price:            decimal.NewFromFloat(12.25),
			},
			{
				ShortDescription: "Knorr Creamy Chicken",
				Price:            decimal.NewFromFloat(1.26),
			},
			{
				ShortDescription: "Doritos Nacho Cheese",
				Price:            decimal.NewFromFloat(3.35),
			},
			{
				ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
				Price:            decimal.NewFromFloat(12.00),
			},
		},
	}
}

// getTestReceipt2 returns a test receipt using the examples in the README.md.
func getTestReceipt2() entity.Receipt {
	return entity.Receipt{
		Retailer:     "M&M Corner Market",
		PurchaseDate: time.Date(2022, 3, 20, 0, 0, 0, 0, time.UTC),
		PurchaseTime: time.Date(0, 1, 1, 14, 33, 0, 0, time.UTC),
		Total:        decimal.NewFromFloat(9.0),
		Items: []entity.Item{
			{
				ShortDescription: "Gatorade",
				Price:            decimal.NewFromFloat(2.25),
			},
			{
				ShortDescription: "Gatorade",
				Price:            decimal.NewFromFloat(2.25),
			},
			{
				ShortDescription: "Gatorade",
				Price:            decimal.NewFromFloat(2.25),
			},
			{
				ShortDescription: "Gatorade",
				Price:            decimal.NewFromFloat(2.25),
			},
		},
	}
}

func TestProcess(t *testing.T) {
	t.Parallel()

	useCase, repo := receiptTestService(t)
	testReceipt1 := getTestReceipt1()
	testReceipt2 := getTestReceipt2()

	tests := []test{
		{
			name: "insert receipt 1",
			mock: func() {
				repo.EXPECT().InsertReceipt(context.Background(), testReceipt1).Return(uint64(1))
			},
			receipt: testReceipt1,
			res:     uint64(1),
		},
		{
			name: "insert receipt 2",
			mock: func() {
				repo.EXPECT().InsertReceipt(context.Background(), testReceipt2).Return(uint64(2))
			},
			receipt: testReceipt2,
			res:     uint64(2),
		},
	}

	for idx := range tests {
		tc := tests[idx]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res := useCase.Process(context.Background(), tc.receipt)

			require.Equal(t, res, tc.res)
		})
	}
}

func TestGetPoints(t *testing.T) {
	t.Parallel()

	useCase, repo := receiptTestService(t)
	testReceipt1 := getTestReceipt1()
	testReceipt2 := getTestReceipt2()

	tests := []test{
		{
			id:   1,
			name: "get receipt 1 points",
			mock: func() {
				repo.EXPECT().GetReceipt(context.Background(), uint64(1)).Return(testReceipt1, nil)
			},
			receipt: testReceipt1,
			res:     28,
			err:     nil,
		},
		{
			id:   2,
			name: "get receipt 2 points",
			mock: func() {
				repo.EXPECT().GetReceipt(context.Background(), uint64(2)).Return(testReceipt2, nil)
			},
			receipt: testReceipt2,
			res:     109,
			err:     nil,
		},
	}

	for idx := range tests {
		tc := tests[idx]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := useCase.GetPoints(context.Background(), tc.id)

			require.Equal(t, tc.res, res)
			require.Equal(t, tc.err, err)
		})
	}
}
