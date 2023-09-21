package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/adriansabvr/receipt_processor/config"
	v1 "github.com/adriansabvr/receipt_processor/internal/controller/http/v1"
	"github.com/adriansabvr/receipt_processor/pkg/httpserver"
	"github.com/adriansabvr/receipt_processor/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/palantir/stacktrace"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting interrupt signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		l.Error(stacktrace.Propagate(err, "failed to start http server"))
	}

	// Shutdown
	err := httpServer.Shutdown()
	if err != nil {
		l.Error(stacktrace.Propagate(err, "failed to shutdown http server"))
	}
}
