package main

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

var logger *zap.Logger

type healthResponse struct {
	OK     bool   `json:"ok"`
	Status string `json:"status"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	resp := healthResponse{OK: true, Status: "up"}

	b, err := json.Marshal(resp)
	if err != nil {
		logger.Error("falha ao fazer marshal do health response", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logger.Info("health check",
		zap.String("path", r.URL.Path),
		zap.Any("body", json.RawMessage(b)),
	)

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func main() {
	logger, _ = zap.NewDevelopment()
	defer logger.Sync()

	http.HandleFunc("GET /health", healthHandler)

	logger.Info("startando API", zap.String("addr", ":8080"))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Fatal("servidor caiu", zap.Error(err))
	}
}
