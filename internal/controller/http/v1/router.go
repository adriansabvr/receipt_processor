package v1

import (
	"github.com/adriansabvr/receipt_processor/internal/usecase/receipt"
	"github.com/adriansabvr/receipt_processor/internal/usecase/receipt/repo"
	"github.com/adriansabvr/receipt_processor/pkg/logger"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, l logger.Interface) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Routers
	h := handler.Group("/v1")

	// ReceiptService
	receiptRepo := repo.New()
	receiptUseCase := receipt.New(receiptRepo)
	newReceiptRoutes(h, receiptUseCase, l)
}
