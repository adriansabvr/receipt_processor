package v1

import (
	"github.com/adriansabvr/receipt_processor/internal/usecase/receipt"
	"github.com/adriansabvr/receipt_processor/internal/usecase/receipt/repo"
	"github.com/adriansabvr/receipt_processor/pkg/logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewRouter -.
//
//	@title			Receipt Processor API
//	@description	This is a sample server for a receipt processor challenge.
//	@version		1.0
//	@host			localhost:8080
//	@BasePath		/v1
func NewRouter(handler *gin.Engine, l logger.Interface) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler) // Routers
	h := handler.Group("/v1")

	// ReceiptService
	receiptRepo := repo.New()
	receiptUseCase := receipt.New(receiptRepo)
	newReceiptRoutes(h, receiptUseCase, l)
}
