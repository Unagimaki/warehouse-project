package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"service-api/internal/config"
	"service-api/internal/handler"
	"time"
)

func Run() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handler.Health)

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: mux,
	}
	quit := make(chan os.Signal, 1)
	errCh := make(chan error, 1)
	signal.Notify(quit, os.Interrupt)
	log.Printf("Starting HTTP server on :%s", cfg.Port)
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			errCh <- err
		}
	}()
	select {
	case err := <-errCh:
		return err
	case <-quit:
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		return err
	}

	log.Println("Server exiting")

	return nil
}
