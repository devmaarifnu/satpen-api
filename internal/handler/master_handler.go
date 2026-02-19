package handler

import (
	"net/http"
	"satpen-api/internal/service"
	"satpen-api/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type MasterHandler struct {
	service service.MasterService
	log     *logrus.Logger
}

func NewMasterHandler(service service.MasterService, log *logrus.Logger) *MasterHandler {
	return &MasterHandler{
		service: service,
		log:     log,
	}
}

// GetAllProvinsi godoc
// @Summary Get all provinsi
// @Description Get list of all provinsi
// @Tags master
// @Accept json
// @Produce json
// @Param search query string false "Search by nama provinsi"
// @Success 200 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /provinsi [get]
func (h *MasterHandler) GetAllProvinsi(c *gin.Context) {
	search := c.Query("search")

	provinsi, err := h.service.GetAllProvinsi(search)
	if err != nil {
		h.log.WithError(err).Error("Failed to get provinsi")
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get provinsi", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Provinsi retrieved successfully", provinsi)
}

// GetProvinsiByID godoc
// @Summary Get provinsi by ID
// @Description Get single provinsi by ID
// @Tags master
// @Accept json
// @Produce json
// @Param id path int true "Provinsi ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /provinsi/{id} [get]
func (h *MasterHandler) GetProvinsiByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID", err.Error())
		return
	}

	provinsi, err := h.service.GetProvinsiByID(uint(id))
	if err != nil {
		h.log.WithError(err).Error("Failed to get provinsi")
		utils.ErrorResponse(c, http.StatusNotFound, "Provinsi not found", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Provinsi retrieved successfully", provinsi)
}

// GetAllKabupaten godoc
// @Summary Get all kabupaten
// @Description Get list of all kabupaten/kota
// @Tags master
// @Accept json
// @Produce json
// @Param provinsi_id query int false "Filter by provinsi ID"
// @Param search query string false "Search by nama kabupaten"
// @Success 200 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /kabupaten [get]
func (h *MasterHandler) GetAllKabupaten(c *gin.Context) {
	var provinsiID uint
	if provinsiIDStr := c.Query("provinsi_id"); provinsiIDStr != "" {
		id, err := strconv.ParseUint(provinsiIDStr, 10, 64)
		if err != nil {
			utils.ErrorResponse(c, http.StatusBadRequest, "Invalid provinsi_id", err.Error())
			return
		}
		provinsiID = uint(id)
	}

	search := c.Query("search")

	kabupaten, err := h.service.GetAllKabupaten(provinsiID, search)
	if err != nil {
		h.log.WithError(err).Error("Failed to get kabupaten")
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get kabupaten", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Kabupaten retrieved successfully", kabupaten)
}

// GetKabupatenByID godoc
// @Summary Get kabupaten by ID
// @Description Get single kabupaten by ID
// @Tags master
// @Accept json
// @Produce json
// @Param id path int true "Kabupaten ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /kabupaten/{id} [get]
func (h *MasterHandler) GetKabupatenByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID", err.Error())
		return
	}

	kabupaten, err := h.service.GetKabupatenByID(uint(id))
	if err != nil {
		h.log.WithError(err).Error("Failed to get kabupaten")
		utils.ErrorResponse(c, http.StatusNotFound, "Kabupaten not found", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Kabupaten retrieved successfully", kabupaten)
}

// GetAllPengurusCabang godoc
// @Summary Get all pengurus cabang
// @Description Get list of all pengurus cabang with pagination
// @Tags master
// @Accept json
// @Produce json
// @Param provinsi_id query int false "Filter by provinsi ID"
// @Param search query string false "Search by nama cabang"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(20)
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /pengurus-cabang [get]
func (h *MasterHandler) GetAllPengurusCabang(c *gin.Context) {
	// Parse pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}

	// Build filters
	filters := make(map[string]interface{})

	if provinsiIDStr := c.Query("provinsi_id"); provinsiIDStr != "" {
		id, err := strconv.ParseUint(provinsiIDStr, 10, 64)
		if err != nil {
			utils.ErrorResponse(c, http.StatusBadRequest, "Invalid provinsi_id", err.Error())
			return
		}
		filters["provinsi_id"] = uint(id)
	}

	if search := c.Query("search"); search != "" {
		filters["search"] = search
	}

	pengurusCabang, total, err := h.service.GetAllPengurusCabang(filters, page, limit)
	if err != nil {
		h.log.WithError(err).Error("Failed to get pengurus cabang")
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get pengurus cabang", err.Error())
		return
	}

	// Calculate pagination
	totalPages := (int(total) + limit - 1) / limit
	hasNext := page < totalPages
	hasPrev := page > 1

	response := map[string]interface{}{
		"pengurus_cabang": pengurusCabang,
		"pagination": map[string]interface{}{
			"current_page":    page,
			"total_pages":     totalPages,
			"total_items":     total,
			"items_per_page":  limit,
			"has_next":        hasNext,
			"has_prev":        hasPrev,
		},
	}

	utils.SuccessResponse(c, http.StatusOK, "Pengurus cabang retrieved successfully", response)
}

// GetPengurusCabangByID godoc
// @Summary Get pengurus cabang by ID
// @Description Get single pengurus cabang by ID
// @Tags master
// @Accept json
// @Produce json
// @Param id path int true "Pengurus Cabang ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /pengurus-cabang/{id} [get]
func (h *MasterHandler) GetPengurusCabangByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID", err.Error())
		return
	}

	pengurusCabang, err := h.service.GetPengurusCabangByID(uint(id))
	if err != nil {
		h.log.WithError(err).Error("Failed to get pengurus cabang")
		utils.ErrorResponse(c, http.StatusNotFound, "Pengurus cabang not found", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Pengurus cabang retrieved successfully", pengurusCabang)
}

// GetAllJenjangPendidikan godoc
// @Summary Get all jenjang pendidikan
// @Description Get list of all jenjang pendidikan
// @Tags master
// @Accept json
// @Produce json
// @Param search query string false "Search by nama jenjang"
// @Success 200 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /jenjang-pendidikan [get]
func (h *MasterHandler) GetAllJenjangPendidikan(c *gin.Context) {
	search := c.Query("search")

	jenjang, err := h.service.GetAllJenjangPendidikan(search)
	if err != nil {
		h.log.WithError(err).Error("Failed to get jenjang pendidikan")
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get jenjang pendidikan", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Jenjang pendidikan retrieved successfully", jenjang)
}

// GetJenjangPendidikanByID godoc
// @Summary Get jenjang pendidikan by ID
// @Description Get single jenjang pendidikan by ID
// @Tags master
// @Accept json
// @Produce json
// @Param id path int true "Jenjang Pendidikan ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /jenjang-pendidikan/{id} [get]
func (h *MasterHandler) GetJenjangPendidikanByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID", err.Error())
		return
	}

	jenjang, err := h.service.GetJenjangPendidikanByID(uint(id))
	if err != nil {
		h.log.WithError(err).Error("Failed to get jenjang pendidikan")
		utils.ErrorResponse(c, http.StatusNotFound, "Jenjang pendidikan not found", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Jenjang pendidikan retrieved successfully", jenjang)
}
