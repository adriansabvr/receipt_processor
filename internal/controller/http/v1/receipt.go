package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/adriansabvr/receipt_processor/internal/entity"
	"github.com/adriansabvr/receipt_processor/internal/usecase/receipt"
	"github.com/adriansabvr/receipt_processor/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/palantir/stacktrace"
	"github.com/shopspring/decimal"
)

type receiptRoutes struct {
	uc     receipt.UseCaseContract
	logger logger.Interface
}

func newReceiptRoutes(handler *gin.RouterGroup, uc receipt.UseCaseContract, logger logger.Interface) {
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

// @Summary		Process receipt
// @Description	Process receipt and return receipt id
// @ID				process
// @Tags			receipt
// @Accept			json
// @Produce		json
// @Param			request	body		processRequest	true	"Set up receipt to process"
// @Success		200		{object}	processResponse
// @Failure		400		{object}	response
// @Router			/receipts/process [post]
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

	items := make([]entity.Item, len(request.Items))
	for i, item := range request.Items {
		items[i] = entity.Item{
			ShortDescription: item.ShortDescription,
			Price:            item.Price,
		}
	}

	receiptEnt := entity.Receipt{
		Retailer:     request.Retailer,
		PurchaseDate: purchaseDate,
		PurchaseTime: purchaseTime,
		Total:        request.Total,
		Items:        items,
	}

	receiptID := r.uc.Process(c.Request.Context(), receiptEnt)

	c.JSON(http.StatusOK, processResponse{ID: receiptID})
}

type pointsResponse struct {
	Points int `json:"points"`
}

// @Summary		Get points
// @Description	Get receipt points by receipt id
// @ID				get-points
// @Tags			receipt
// @Accept			json
// @Produce		json
// @Success		200	{object}	pointsResponse
// @Failure		400	{object}	response
// @Failure		500	{object}	response
// @Router			/receipts/:id/points [get]
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
