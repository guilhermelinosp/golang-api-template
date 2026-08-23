// Package main provides a minimal HTTP API entry point for golang-api-template.
// Replace handlers with your actual domain routes and middleware.
package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/guilhermelinosp/hellnet-lib-telemetry/telemetry"
)

func main() {
	ops, err := telemetry.New(telemetry.Options{
		ServiceName: "golang-api-template",
		Enabled:     os.Getenv("HELLNET_TELEMETRY_ENABLED") == "true",
	})
	if err != nil {
		log.Fatalf("failed to init telemetry: %v", err)
	}
	defer func() { _ = ops.Shutdown() }()

	mux := http.NewServeMux()
	mux.Handle("GET /live", ops.Live())
	mux.Handle("GET /ready", ops.Ready())
	mux.Handle("GET /health", ops.Health())
	mux.Handle("GET /metrics", ops.MetricsHandler())
	mux.HandleFunc("/", rootHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      telemetry.Middleware(ops, mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		ops.Log().Info("API listening", "addr", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			ops.Log().Error("server error", "error", err)
			os.Exit(1)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	ops.Log().Info("shutting down gracefully")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		ops.Log().Error("shutdown error", "error", err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("golang-api-template\n"))
}
