package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/dwaynedwards/sell-u-lar/pkg/db"
	"github.com/dwaynedwards/sell-u-lar/products"
	"github.com/dwaynedwards/sell-u-lar/products/internal/grpc"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	db := db.NewPostgres(products.Config.DatabaseURL)
	s := grpc.NewServer(db)

	slog.Info("Server starting up")

	if err := s.Start(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		slog.Error("Sever starting up failed", slog.Any("err", err))
		err = s.Stop()
		if err != nil {
			slog.Error("Sever shutting down failed", slog.Any("err", err))
		}
		os.Exit(1)
	}

	<-sig

	slog.Info("Server shutting down")

	if err := s.Stop(); err != nil {
		slog.Error("Sever shutting down failed", slog.Any("err", err))
		os.Exit(1)
	}

	slog.Info("Server shut down")
}
