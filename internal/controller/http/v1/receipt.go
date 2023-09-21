package v1

import (
	"github.com/adriansabvr/receipt_processor/internal/entity"
	"github.com/adriansabvr/receipt_processor/internal/usecase"
	"github.com/adriansabvr/receipt_processor/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/palantir/stacktrace"
	"github.com/shopspring/decimal"
	"net/http"
	"strconv"
	"time"
)

type receiptRoutes struct {
	uc     usecase.Receipt
	logger logger.Interface
}

func newReceiptRoutes(handler *gin.RouterGroup, uc usecase.Receipt, logger logger.Interface) {
	r := &receiptRoutes{uc, logger}

	h := handler.Group("/receipts")
	{
		h.GET("/:id/points", r.getPoints)
		h.POST("/process", r.process)
	}
}

type processRequest struct {
	Retailer     string          `json:"retailer"`
	PurchaseDate string          `json:"purchaseDate"`
	PurchaseTime string          `json:"purchaseTime"`
	Total        decimal.Decimal `json:"total"`
	Items        []struct {
		ShortDescription string          `json:"shortDescription"`
		Price            decimal.Decimal `json:"price"`
	} `json:"items"`
}

type processResponse struct {
	ID uint64 `json:"id"`
}

func (r *receiptRoutes) process(c *gin.Context) {

	var request processRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		BadRequest(c, r.logger, err)

		return
	}

	purchaseDate, err := time.Parse(time.DateOnly, request.PurchaseDate)
	if err != nil {
		BadRequest(c, r.logger, stacktrace.Propagate(err, "failed to parse purchase date"))

		return
	}

	purchaseTime, err := time.Parse("15:04", request.PurchaseTime)
	if err != nil {
		BadRequest(c, r.logger, stacktrace.Propagate(err, "failed to parse purchase time"))

		return
	}

	receipt := entity.Receipt{
		Retailer:     request.Retailer,
		PurchaseDate: purchaseDate,
		PurchaseTime: purchaseTime,
		Total:        request.Total,
		Items:        request.Items,
	}

	receiptID, err := r.uc.Process(c.Request.Context(), receipt)
	if err != nil {
		InternalServerError(c, r.logger, stacktrace.Propagate(err, "failed to process receipt"))

		return
	}

	c.JSON(http.StatusOK, processResponse{ID: receiptID})
}

type pointsResponse struct {
	Points int `json:"points"`
}

func (r *receiptRoutes) getPoints(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		BadRequest(c, r.logger, stacktrace.Propagate(err, "failed to parse receipt id"))

		return
	}

	points, err := r.uc.GetPoints(c.Request.Context(), id)
	if err != nil {
		InternalServerError(c, r.logger, stacktrace.Propagate(err, "failed to get points"))

		return
	}

	OK(c, pointsResponse{Points: points})
}
