package router

import (
	"github.com/shortformikael/Hermes/src/internal/logger"
)

func Init() {
	logger.Info("Initializing router...")
	defer logger.Info("Router initialized successfully")
}
