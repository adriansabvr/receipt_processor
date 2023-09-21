package main

import (
	"github.com/adriansabvr/receipt_processor/config"
	"github.com/adriansabvr/receipt_processor/internal/app"
	"log"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
