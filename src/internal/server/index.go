package server

import (
	// "fmt"
	// "net/http"

	// "github.com/shortformikael/Hermes/src/internal/config"
	"github.com/shortformikael/Hermes/src/internal/logger"
)

func Init() {
	logger.Info("Initializing server...")
	defer logger.Info("Server initialized successfully")
}

// func StartServer() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte("Hello, World!"))
// 	})
// 	err := http.ListenAndServe(config.GetConfig().Port, nil)
// 	if err != nil {
// 		logger.Error(fmt.Sprintf("Failed to start server: %v", err))
// 	}
// }
