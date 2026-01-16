package routes

import (
	"satpen-api/internal/config"
	"satpen-api/internal/handler"
	"satpen-api/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SetupRoutes(
	r *gin.Engine,
	cfg *config.Config,
	log *logrus.Logger,
	satpenHandler *handler.SatpenHandler,
	masterHandler *handler.MasterHandler,
	healthHandler *handler.HealthHandler,
) {
	// Middleware
	r.Use(middleware.CORS(cfg))
	r.Use(middleware.Logger(log))
	r.Use(gin.Recovery())

	// Start cleanup routine for rate limiter
	middleware.CleanupVisitors(5 * 60) // 5 minutes

	// Health check
	if cfg.Monitoring.Enabled {
		r.GET(cfg.Monitoring.HealthCheckPath, healthHandler.HealthCheck)
	}

	// API v1 routes
	v1 := r.Group(cfg.API.BasePath)
	{
		// Satpen endpoints
		satpen := v1.Group("/satpen")
		{
			satpen.GET("", middleware.RateLimit(cfg, cfg.RateLimit.Satpen), satpenHandler.GetAllSatpen)
			satpen.GET("/statistics", middleware.RateLimit(cfg, cfg.RateLimit.Statistics), satpenHandler.GetStatistics)
			satpen.GET("/:id", middleware.RateLimit(cfg, cfg.RateLimit.Satpen), satpenHandler.GetSatpenByID)
		}

		// Provinsi endpoints
		provinsi := v1.Group("/provinsi")
		{
			provinsi.GET("", middleware.RateLimit(cfg, cfg.RateLimit.Satpen), masterHandler.GetAllProvinsi)
			provinsi.GET("/:id", middleware.RateLimit(cfg, cfg.RateLimit.Satpen), masterHandler.GetProvinsiByID)
		}

		// Kabupaten endpoints
		kabupaten := v1.Group("/kabupaten")
		{
			kabupaten.GET("", middleware.RateLimit(cfg, cfg.RateLimit.Satpen), masterHandler.GetAllKabupaten)
			kabupaten.GET("/:id", middleware.RateLimit(cfg, cfg.RateLimit.Satpen), masterHandler.GetKabupatenByID)
		}

		// Pengurus Cabang endpoints
		pengurusCabang := v1.Group("/pengurus-cabang")
		{
			pengurusCabang.GET("", middleware.RateLimit(cfg, cfg.RateLimit.Satpen), masterHandler.GetAllPengurusCabang)
			pengurusCabang.GET("/:id", middleware.RateLimit(cfg, cfg.RateLimit.Satpen), masterHandler.GetPengurusCabangByID)
		}

		// Jenjang Pendidikan endpoints
		jenjangPendidikan := v1.Group("/jenjang-pendidikan")
		{
			jenjangPendidikan.GET("", middleware.RateLimit(cfg, cfg.RateLimit.Satpen), masterHandler.GetAllJenjangPendidikan)
			jenjangPendidikan.GET("/:id", middleware.RateLimit(cfg, cfg.RateLimit.Satpen), masterHandler.GetJenjangPendidikanByID)
		}
	}
}
