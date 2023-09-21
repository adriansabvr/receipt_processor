package v1

import (
	"github.com/adriansabvr/receipt_processor/internal/usecase"
	"github.com/adriansabvr/receipt_processor/internal/usecase/repo"
	"github.com/adriansabvr/receipt_processor/pkg/logger"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, l logger.Interface) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Routers
	h := handler.Group("/v1")

	// Receipt
	receiptRepo := repo.New()
	receiptUseCase := usecase.New(receiptRepo)
	newReceiptRoutes(h, receiptUseCase, l)
}
