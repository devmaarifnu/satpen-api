package main

import (
	"fmt"
	"log"
	"os"
	"satpen-api/internal/config"
	"satpen-api/internal/database"
	"satpen-api/internal/handler"
	"satpen-api/internal/repository"
	"satpen-api/internal/routes"
	"satpen-api/internal/service"

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

	// Start server
	addr := fmt.Sprintf(":%d", cfg.App.Port)
	logger.Infof("Server starting on %s", addr)
	logger.Infof("API available at http://localhost%s%s", addr, cfg.API.BasePath)
	logger.Infof("Health check available at http://localhost%s%s", addr, cfg.Monitoring.HealthCheckPath)

	if err := r.Run(addr); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
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
