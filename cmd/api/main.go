package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"simple-bank/internal/app"
	"simple-bank/internal/config"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg := config.Load()

	if err := app.Run(ctx, cfg); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
