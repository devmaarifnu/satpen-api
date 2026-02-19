package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"satpen-api/internal/config"
	"satpen-api/internal/database"
	"satpen-api/internal/handler"
	"satpen-api/internal/middleware"
	"satpen-api/internal/repository"
	"satpen-api/internal/routes"
	"satpen-api/internal/service"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Setup logger
	logger := setupLogger(cfg)
	logger.Infof("Starting %s v%s [%s]", cfg.App.Name, cfg.App.Version, cfg.App.Env)

	// Connect to database
	db, err := database.Connect(cfg)
	if err != nil {
		logger.Fatalf("Failed to connect to database: %v", err)
	}
	logger.Info("Database connected successfully")

	// Auto migrate
	if err := database.AutoMigrate(db); err != nil {
		logger.Fatalf("Failed to run auto migration: %v", err)
	}

	// Initialize repositories
	satpenRepo := repository.NewSatpenRepository(db)
	masterRepo := repository.NewMasterRepository(db)

	// Initialize services
	satpenService := service.NewSatpenService(satpenRepo, cfg)
	masterService := service.NewMasterService(masterRepo)

	// Initialize handlers
	satpenHandler := handler.NewSatpenHandler(satpenService)
	masterHandler := handler.NewMasterHandler(masterService, logger)
	healthHandler := handler.NewHealthHandler(cfg)

	// Setup Gin
	if cfg.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// Setup routes
	routes.SetupRoutes(r, cfg, logger, satpenHandler, masterHandler, healthHandler)

	// Create context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start cleanup routine for rate limiter (ONCE)
	if cfg.RateLimit.Enabled {
		middleware.StartCleanup(ctx, 5*time.Minute)
		logger.Info("Rate limiter cleanup routine started")
	}

	// Create HTTP server
	addr := fmt.Sprintf(":%d", cfg.App.Port)
	srv := &http.Server{
		Addr:           addr,
		Handler:        r,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB
	}

	// Start server in goroutine
	go func() {
		logger.Infof("Server starting on %s", addr)
		logger.Infof("API available at http://localhost%s%s", addr, cfg.API.BasePath)
		logger.Infof("Health check available at http://localhost%s%s", addr, cfg.Monitoring.HealthCheckPath)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	// Cancel context to stop cleanup routine
	cancel()
	middleware.StopCleanup()

	// Graceful shutdown with timeout
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Errorf("Server forced to shutdown: %v", err)
	}

	// Close database connection
	if err := database.Close(); err != nil {
		logger.Errorf("Error closing database: %v", err)
	}

	logger.Info("Server exited gracefully")
}

func setupLogger(cfg *config.Config) *logrus.Logger {
	logger := logrus.New()

	// Set log level
	switch cfg.Logging.Level {
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	default:
		logger.SetLevel(logrus.InfoLevel)
	}

	// Set log format
	if cfg.Logging.Format == "json" {
		logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}

	// Set output
	if cfg.Logging.Output == "file" {
		file, err := os.OpenFile(cfg.Logging.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			logger.SetOutput(file)
		} else {
			log.Printf("Failed to log to file, using default stdout: %v", err)
		}
	} else {
		logger.SetOutput(os.Stdout)
	}

	return logger
}
