package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"service-werehouse/internal/config"
	"service-werehouse/internal/database"
	"service-werehouse/internal/handler"
	"service-werehouse/internal/middleware"
	"service-werehouse/internal/repository"
	"service-werehouse/internal/service"
	"time"
)

func Run() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handler.Health)

	dbCfg := database.DBConfig{
		Driver:   cfg.DB.Driver,
		Host:     cfg.DB.Host,
		Port:     cfg.DB.PortNum,
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
		DBName:   cfg.DB.DBName,
		SSLMode:  cfg.DB.SSLMode,
	}

	if err := database.RunMigrations(dbCfg); err != nil {
		return err
	}

	db, err := database.NewDB(dbCfg)
	if err != nil {
		return err
	}
	defer db.Close()

	repo := repository.NewProductRepository(db)
	productService := service.NewProductService(repo)
	productHandler := handler.NewProductHandler(productService)
	mux.HandleFunc("/products", productHandler.CreateProduct)

	muxhandler := middleware.LoggingMiddleware(mux)

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: muxhandler,
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
