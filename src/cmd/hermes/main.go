// Package hermes: Execute the program
package hermes

import (
	"github.com/shortformikael/Hermes/src/internal/config"
	"github.com/shortformikael/Hermes/src/internal/logger"
	"github.com/shortformikael/Hermes/src/internal/router"
	"github.com/shortformikael/Hermes/src/internal/server"
)

func Init() {
	logger.Info("Initializing...")
	defer logger.Info("Hermes initialized successfully")
	config.Init()
	server.Init()
	router.Init()
}

func Exec() {
	logger.Info("Executing...")
	defer logger.Info("Hermes executed successfully")
}
