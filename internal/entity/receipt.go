package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

// Receipt -.
type Receipt struct {
	ID           uint64          `json:"id"`
	Retailer     string          `json:"retailer"`
	PurchaseDate time.Time       `json:"purchaseDate"`
	PurchaseTime time.Time       `json:"purchaseTime"`
	Total        decimal.Decimal `json:"total"`
	Items        []Item          `json:"items"`
}

type Item struct {
	ShortDescription string          `json:"shortDescription"`
	Price            decimal.Decimal `json:"price"`
}
