package app

import (
	"github.com/adriansabvr/receipt_processor/config"
	"github.com/adriansabvr/receipt_processor/pkg/httpserver"
	"github.com/adriansabvr/receipt_processor/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/palantir/stacktrace"
	"os"
	"os/signal"
	"syscall"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// HTTP Server
	handler := gin.New()
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting interrupt signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		l.Error(stacktrace.Propagate(err, "app - Run - httpServer.Notify"))
	}

	// Shutdown
	err := httpServer.Shutdown()
	if err != nil {
		l.Error(stacktrace.Propagate(err, "app - Run - httpServer.Shutdown"))
	}
}