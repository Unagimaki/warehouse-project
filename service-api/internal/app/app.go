package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"service-api/internal/client/warehouse"
	"service-api/internal/config"
	"service-api/internal/handler"
	"service-api/internal/metrics"
	"service-api/internal/middleware"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Run() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}
	metrics.Init()
	warehouseClient := warehouse.New(cfg.WarehouseURL)
	productHandler := handler.New(warehouseClient)

	mux := http.NewServeMux()

	mux.HandleFunc("/health", handler.Health)
	mux.HandleFunc("POST /products", productHandler.CreateProduct)
	mux.Handle("/metrics", promhttp.Handler())

	handler := metrics.Middleware(mux)
	handler = middleware.Logging(handler)

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: handler,
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
