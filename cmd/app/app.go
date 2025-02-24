package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	handler "github.com/ruziba3vich/music_lib/internal/http"
	"github.com/ruziba3vich/music_lib/internal/service"
	"github.com/ruziba3vich/music_lib/internal/storage"
	"github.com/ruziba3vich/music_lib/pkg/config"
)

// Run initializes and starts the application with graceful shutdown
func Run(logger *log.Logger) error {
	// Load configuration
	cfg := config.LoadConfig()
	// Connect to database
	db, err := storage.GetDBConnection(cfg)
	if err != nil {
		logger.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize storage
	store := storage.NewStorage(db)

	// Initialize service layer
	service := service.NewService(store, logger)

	// Initialize handler layer
	handler := handler.NewHandler(service, logger)

	// Initialize Gin router
	router := gin.Default()

	// Set up routes
	handler.RegisterRoutes(router)

	// Start the server
	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	// Run server in a goroutine
	go func() {
		logger.Printf("Starting server on port %s", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Printf("Server error: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit // Wait for termination signal
	logger.Println("Shutting down server...")

	// Create a context with a timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Printf("Server shutdown error: %v", err)
		return err
	}

	logger.Println("Server shutdown gracefully")
	return nil
}
