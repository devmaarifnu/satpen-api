package handler

import (
	"net/http"
	"satpen-api/internal/config"
	"satpen-api/internal/database"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	cfg *config.Config
}

func NewHealthHandler(cfg *config.Config) *HealthHandler {
	return &HealthHandler{cfg: cfg}
}

// HealthCheck handles GET /health
func (h *HealthHandler) HealthCheck(c *gin.Context) {
	// Check database connection
	sqlDB, err := database.DB.DB()
	dbStatus := "healthy"
	if err != nil || sqlDB.Ping() != nil {
		dbStatus = "unhealthy"
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "ok",
		"app":      h.cfg.App.Name,
		"version":  h.cfg.App.Version,
		"database": dbStatus,
	})
}
