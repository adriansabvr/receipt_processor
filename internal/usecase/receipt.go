package usecase

import (
	"context"
	"github.com/adriansabvr/receipt_processor/internal/entity"
	"github.com/palantir/stacktrace"
	"github.com/shopspring/decimal"
	"strings"
	"unicode"
)

// ReceiptUseCase -.
type ReceiptUseCase struct {
	repo ReceiptRepo
}

// New -.
func New(r ReceiptRepo) *ReceiptUseCase {
	return &ReceiptUseCase{
		repo: r,
	}
}

// Process - inserts receipt to repo.
func (uc *ReceiptUseCase) Process(ctx context.Context, receipt entity.Receipt) (uint64, error) {
	receiptID, err := uc.repo.InsertReceipt(ctx, receipt)
	if err != nil {
		return 0, stacktrace.Propagate(err, "failed to insert receipt to repo")
	}

	return receiptID, nil
}

// GetPoints -.
func (uc *ReceiptUseCase) GetPoints(ctx context.Context, receiptID uint64) (int, error) {
	receipt, err := uc.repo.GetReceipt(ctx, receiptID)
	if err != nil {
		return 0, stacktrace.Propagate(err, "failed to get receipt from repo")
	}

	points, err := getPoints(receipt)
	if err != nil {
		return 0, stacktrace.Propagate(err, "failed to get points")
	}

	return points, nil
}

func getPoints(receipt entity.Receipt) (int, error) {
	points := 0

	// One point for every alphanumeric character in the retailer name.
	points += countAlphanumeric(receipt.Retailer)

	// 50 points if the total is a round dollar amount with no cents.
	if receipt.Total.IsInteger() {
		points += 50
	}

	// 25 points if the total is a multiple of 0.25.
	if receipt.Total.Mod(decimal.NewFromFloat(0.25)).Equal(decimal.Zero) {
		points += 25
	}

	// 5 points for every two items on the receipt.
	points += len(receipt.Items) / 2 * 5

	// If the trimmed length of the item description is a multiple of 3
	// multiply the price by 0.2 and round up to the nearest integer.
	// The result is the number of points earned.
	for _, item := range receipt.Items {
		trimmedDescription := strings.Trim(item.ShortDescription, " ")
		if len(trimmedDescription)%3 == 0 {
			extraPoints := item.Price.Mul(decimal.NewFromFloat(0.2)).Ceil()
			points += int(extraPoints.IntPart())
		}
	}

	// 6 points if the day in the purchase date is odd.
	if receipt.PurchaseDate.Day()%2 == 1 {
		points += 6
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	if (receipt.PurchaseTime.Hour() > 13 && receipt.PurchaseTime.Minute() > 1) && receipt.PurchaseTime.Hour() < 16 {
		points += 10
	}

	return points, nil
}

func countAlphanumeric(str string) int {
	count := 0
	for _, r := range str {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			count++
		}
	}

	return count
}
