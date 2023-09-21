package receipt

import (
	"context"
	"strings"
	"unicode"

	"github.com/adriansabvr/receipt_processor/internal/entity"
	"github.com/palantir/stacktrace"
	"github.com/shopspring/decimal"
)

// UseCase -.
type UseCase struct {
	repo RepoContract
}

// New -.
func New(r RepoContract) *UseCase {
	return &UseCase{
		repo: r,
	}
}

// Process - inserts receipt to repo.
func (uc *UseCase) Process(ctx context.Context, receipt entity.Receipt) uint64 {
	receiptID := uc.repo.InsertReceipt(ctx, receipt)

	return receiptID
}

// GetPoints -.
func (uc *UseCase) GetPoints(ctx context.Context, receiptID uint64) (int, error) {
	receipt, err := uc.repo.GetReceipt(ctx, receiptID)
	if err != nil {
		return 0, stacktrace.Propagate(err, "failed to get receipt from repo")
	}

	points := getPoints(receipt)

	return points, nil
}

func getPoints(receipt entity.Receipt) int {
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

	return points
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
