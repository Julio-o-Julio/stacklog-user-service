package main

import (
	"github.com/julio-o-julio/stacklog-user-service/config"
	"github.com/julio-o-julio/stacklog-user-service/internal/handler"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize configs
	err := config.Initialize()
	if err != nil {
		logger.Errorf("| pkg main | Config initialization error: %v", err)
		return
	}

	// Initialize handler
	handler.Initialize()
}