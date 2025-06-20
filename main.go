package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/madsfranzen/go-webserver/database"
)

func main() {

	// Logging
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Database
	if err := database.Connect(); err != nil {
		slog.Error("DB connection failed", "error", err)
		os.Exit(1)
	}

	defer database.Close()

	// Routes
	r := setupRouter()

	slog.Info("Server running on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		slog.Error("Server error", "error", err)
	}
}
