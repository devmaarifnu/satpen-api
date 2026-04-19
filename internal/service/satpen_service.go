package service

import (
	"bytes"
	"errors"
	"fmt"
	"satpen-api/internal/config"
	"satpen-api/internal/models"
	"satpen-api/internal/repository"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type SatpenService interface {
	GetAllSatpen(filters map[string]interface{}, page, limit int, sort string, includeStats bool) ([]models.Satpen, *PaginationMeta, *models.SatpenStatistics, error)
	GetSatpenByID(id string) (*models.Satpen, error)
	GetStatistics(filters map[string]interface{}) (*models.SatpenStatistics, error)
	ExportSatpen(filters map[string]interface{}, sort string) (*bytes.Buffer, string, error)
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

func (s *satpenService) GetAllSatpen(filters map[string]interface{}, page, limit int, sort string, includeStats bool) ([]models.Satpen, *PaginationMeta, *models.SatpenStatistics, error) {
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

	// Get statistics only if requested (performance optimization)
	var stats *models.SatpenStatistics
	if includeStats {
		stats, err = s.repo.GetStatistics(filters)
		if err != nil {
			return nil, nil, nil, err
		}
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

func (s *satpenService) ExportSatpen(filters map[string]interface{}, sort string) (*bytes.Buffer, string, error) {
	satpenList, err := s.repo.FindAllForExport(filters, sort)
	if err != nil {
		return nil, "", err
	}

	f := excelize.NewFile()
	defer f.Close()

	sheet := "Data Satpen"
	f.SetSheetName("Sheet1", sheet)

	// Header style
	headerStyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Color: "FFFFFF", Size: 11},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"1F4E79"}, Pattern: 1},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center", WrapText: true},
		Border: []excelize.Border{
			{Type: "left", Color: "FFFFFF", Style: 1},
			{Type: "right", Color: "FFFFFF", Style: 1},
			{Type: "top", Color: "FFFFFF", Style: 1},
			{Type: "bottom", Color: "FFFFFF", Style: 1},
		},
	})

	// Data style (odd rows)
	dataStyle, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Vertical: "center", WrapText: true},
		Border: []excelize.Border{
			{Type: "left", Color: "D0D0D0", Style: 1},
			{Type: "right", Color: "D0D0D0", Style: 1},
			{Type: "top", Color: "D0D0D0", Style: 1},
			{Type: "bottom", Color: "D0D0D0", Style: 1},
		},
	})

	// Data style (even rows)
	dataStyleAlt, _ := f.NewStyle(&excelize.Style{
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"EBF3FB"}, Pattern: 1},
		Alignment: &excelize.Alignment{Vertical: "center", WrapText: true},
		Border: []excelize.Border{
			{Type: "left", Color: "D0D0D0", Style: 1},
			{Type: "right", Color: "D0D0D0", Style: 1},
			{Type: "top", Color: "D0D0D0", Style: 1},
			{Type: "bottom", Color: "D0D0D0", Style: 1},
		},
	})

	headers := []string{
		"No", "NPSN", "No. Registrasi", "Nama Satuan Pendidikan",
		"Jenjang", "Provinsi", "Kabupaten", "Kecamatan", "Kelurahan", "Alamat",
		"Kepala Sekolah", "Yayasan", "Tahun Berdiri", "Status", "Akreditasi",
		"Jumlah Siswa", "Jumlah Guru",
	}
	colWidths := []float64{5, 12, 18, 40, 8, 25, 25, 20, 20, 40, 25, 30, 14, 16, 12, 14, 12}

	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
		f.SetCellStyle(sheet, cell, cell, headerStyle)
		f.SetColWidth(sheet, colName(i+1), colName(i+1), colWidths[i])
	}
	f.SetRowHeight(sheet, 1, 30)

	for rowIdx, s := range satpenList {
		row := rowIdx + 2
		style := dataStyle
		if rowIdx%2 == 1 {
			style = dataStyleAlt
		}

		jenjang := ""
		if s.Jenjang != nil {
			jenjang = s.Jenjang.NmJenjang
		}
		provinsi := ""
		if s.Provinsi != nil {
			provinsi = s.Provinsi.NmProv
		}
		kabupaten := ""
		if s.Kabupaten != nil {
			kabupaten = s.Kabupaten.NamaKab
		}
		akreditasi := s.Akreditasi
		if akreditasi == "" {
			akreditasi = "-"
		}

		values := []interface{}{
			rowIdx + 1,
			s.NPSN,
			s.NoRegistrasi,
			s.NmSatpen,
			jenjang,
			provinsi,
			kabupaten,
			s.Kecamatan,
			s.Kelurahan,
			s.Alamat,
			s.Kepsek,
			s.Yayasan,
			s.ThnBerdiri,
			s.Status,
			akreditasi,
			s.JumlahSiswa,
			s.JumlahGuru,
		}

		for colIdx, val := range values {
			cell, _ := excelize.CoordinatesToCellName(colIdx+1, row)
			f.SetCellValue(sheet, cell, val)
			f.SetCellStyle(sheet, cell, cell, style)
		}
		f.SetRowHeight(sheet, row, 20)
	}

	// Freeze the header row
	f.SetPanes(sheet, &excelize.Panes{
		Freeze:      true,
		Split:       false,
		YSplit:      1,
		TopLeftCell: "A2",
		ActivePane:  "bottomLeft",
	})

	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, "", err
	}

	filename := fmt.Sprintf("data-satpen-%s.xlsx", time.Now().Format("20060102-150405"))
	return buf, filename, nil
}

// colName converts a 1-based column index to an Excel column letter (A, B, ..., Z, AA, ...)
func colName(idx int) string {
	name, _ := excelize.ColumnNumberToName(idx)
	return name
}
