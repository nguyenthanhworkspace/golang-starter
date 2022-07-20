package main

import (
	"log"

	"github.com/nguyenthanhworkspace/golang-starter/config"
	"github.com/nguyenthanhworkspace/golang-starter/internal/app"
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
