package handler

import (
	"net/http"
	"satpen-api/internal/service"
	"satpen-api/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SatpenHandler struct {
	service service.SatpenService
}

func NewSatpenHandler(service service.SatpenService) *SatpenHandler {
	return &SatpenHandler{service: service}
}

// GetAllSatpen handles GET /api/v1/satpen
func (h *SatpenHandler) GetAllSatpen(c *gin.Context) {
	// Parse query parameters
	filters := make(map[string]interface{})

	if jenjang := c.Query("jenjang"); jenjang != "" {
		filters["jenjang"] = jenjang
	}

	if provinsi := c.Query("provinsi"); provinsi != "" {
		filters["provinsi"] = provinsi
	}

	if kabupaten := c.Query("kabupaten"); kabupaten != "" {
		filters["kabupaten"] = kabupaten
	}

	if search := c.Query("search"); search != "" {
		filters["search"] = search
	}

	if akreditasi := c.Query("akreditasi"); akreditasi != "" {
		filters["akreditasi"] = akreditasi
	}

	if status := c.Query("status"); status != "" {
		filters["status"] = status
	}

	if verified := c.Query("verified"); verified != "" {
		if verified == "true" {
			filters["verified"] = true
		} else if verified == "false" {
			filters["verified"] = false
		}
	}

	// Parse pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	// Parse sort
	sort := c.DefaultQuery("sort", "-created_at")

	// Check if statistics are needed (skip by default for performance)
	includeStats := c.Query("include_stats") == "true"

	// Get data from service
	satpen, pagination, stats, err := h.service.GetAllSatpen(filters, page, limit, sort, includeStats)
	if err != nil {
		utils.InternalErrorResponse(c, err)
		return
	}

	// Response
	responseData := gin.H{
		"satpen":     satpen,
		"pagination": pagination,
	}

	if includeStats && stats != nil {
		responseData["statistics"] = stats
	}

	utils.SuccessResponse(c, http.StatusOK, "Satuan pendidikan retrieved successfully", responseData)
}

// GetSatpenByID handles GET /api/v1/satpen/:id
func (h *SatpenHandler) GetSatpenByID(c *gin.Context) {
	id := c.Param("id")

	satpen, err := h.service.GetSatpenByID(id)
	if err != nil {
		if err.Error() == "satuan pendidikan not found" {
			utils.NotFoundResponse(c, "Satuan pendidikan not found")
			return
		}
		utils.InternalErrorResponse(c, err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Satuan pendidikan retrieved successfully", satpen)
}

// GetStatistics handles GET /api/v1/satpen/statistics
func (h *SatpenHandler) GetStatistics(c *gin.Context) {
	// Parse filters
	filters := make(map[string]interface{})

	if provinsi := c.Query("provinsi"); provinsi != "" {
		filters["provinsi"] = provinsi
	}

	if jenjang := c.Query("jenjang"); jenjang != "" {
		filters["jenjang"] = jenjang
	}

	// Get statistics
	stats, err := h.service.GetStatistics(filters)
	if err != nil {
		utils.InternalErrorResponse(c, err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Statistics retrieved successfully", stats)
}
