package service

import (
	"errors"
	"satpen-api/internal/config"
	"satpen-api/internal/models"
	"satpen-api/internal/repository"
	"strconv"

	"gorm.io/gorm"
)

type SatpenService interface {
	GetAllSatpen(filters map[string]interface{}, page, limit int, sort string) ([]models.Satpen, *PaginationMeta, *models.SatpenStatistics, error)
	GetSatpenByID(id string) (*models.Satpen, error)
	GetStatistics(filters map[string]interface{}) (*models.SatpenStatistics, error)
}

type satpenService struct {
	repo repository.SatpenRepository
	cfg  *config.Config
}

type PaginationMeta struct {
	CurrentPage  int   `json:"current_page"`
	TotalPages   int   `json:"total_pages"`
	TotalItems   int64 `json:"total_items"`
	ItemsPerPage int   `json:"items_per_page"`
	HasNext      bool  `json:"has_next"`
	HasPrev      bool  `json:"has_prev"`
}

func NewSatpenService(repo repository.SatpenRepository, cfg *config.Config) SatpenService {
	return &satpenService{
		repo: repo,
		cfg:  cfg,
	}
}

func (s *satpenService) GetAllSatpen(filters map[string]interface{}, page, limit int, sort string) ([]models.Satpen, *PaginationMeta, *models.SatpenStatistics, error) {
	// Validate and set defaults for pagination
	if page < 1 {
		page = s.cfg.Pagination.DefaultPage
	}
	if limit < 1 {
		limit = s.cfg.Pagination.DefaultLimit
	}
	if limit > s.cfg.Pagination.MaxLimit {
		limit = s.cfg.Pagination.MaxLimit
	}

	// Get data from repository
	satpen, total, err := s.repo.FindAll(filters, page, limit, sort)
	if err != nil {
		return nil, nil, nil, err
	}

	// Calculate pagination
	totalPages := int(total) / limit
	if int(total)%limit > 0 {
		totalPages++
	}

	pagination := &PaginationMeta{
		CurrentPage:  page,
		TotalPages:   totalPages,
		TotalItems:   total,
		ItemsPerPage: limit,
		HasNext:      page < totalPages,
		HasPrev:      page > 1,
	}

	// Get statistics
	stats, err := s.repo.GetStatistics(filters)
	if err != nil {
		return nil, nil, nil, err
	}

	return satpen, pagination, stats, nil
}

func (s *satpenService) GetSatpenByID(id string) (*models.Satpen, error) {
	// Try to parse as numeric ID first
	if numericID, err := strconv.ParseUint(id, 10, 64); err == nil {
		satpen, err := s.repo.FindByID(uint(numericID))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("satuan pendidikan not found")
			}
			return nil, err
		}
		return satpen, nil
	}

	// Try to find by NPSN
	satpen, err := s.repo.FindByNPSN(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("satuan pendidikan not found")
		}
		return nil, err
	}
	return satpen, nil
}

func (s *satpenService) GetStatistics(filters map[string]interface{}) (*models.SatpenStatistics, error) {
	return s.repo.GetStatistics(filters)
}
