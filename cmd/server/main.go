package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/guilherme0s/IDaaS/pkg/api"
	"github.com/guilherme0s/IDaaS/pkg/config"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, err := config.Init()
	if err != nil {
		slog.Error("log error", "error", err)
	}

	server := api.NewHTTPServer(cfg)

	if err := server.Run(ctx); err != nil {
		slog.Error("server terminated with error", "error", err)
		os.Exit(1)
	}

	slog.Info("application shutdown complete")
}
