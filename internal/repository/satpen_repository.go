package repository

import (
	"satpen-api/internal/models"
	"strings"

	"gorm.io/gorm"
)

type SatpenRepository interface {
	FindAll(filters map[string]interface{}, page, limit int, sort string) ([]models.Satpen, int64, error)
	FindByID(id uint) (*models.Satpen, error)
	FindByNPSN(npsn string) (*models.Satpen, error)
	GetStatistics(filters map[string]interface{}) (*models.SatpenStatistics, error)
	CountByJenjang(filters map[string]interface{}) ([]models.JenjangCount, error)
	CountByAkreditasi(filters map[string]interface{}) ([]models.AkreditasiCount, error)
	GetTopProvinsi(limit int) ([]models.ProvinsiStats, error)
}

type satpenRepository struct {
	db *gorm.DB
}

func NewSatpenRepository(db *gorm.DB) SatpenRepository {
	return &satpenRepository{db: db}
}

func (r *satpenRepository) FindAll(filters map[string]interface{}, page, limit int, sort string) ([]models.Satpen, int64, error) {
	var satpen []models.Satpen
	var total int64

	query := r.db.Model(&models.Satpen{}).
		Preload("Provinsi").
		Preload("Kabupaten").
		Preload("Jenjang").
		Preload("Kategori").
		Preload("PengurusCabang").
		Preload("PDPTK", func(db *gorm.DB) *gorm.DB {
			// Get latest PDPTK data
			return db.Order("tapel DESC").Limit(1)
		})

	// Apply filters
	query = r.applyFilters(query, filters)

	// Count total
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination
	offset := (page - 1) * limit
	query = query.Offset(offset).Limit(limit)

	// Sorting
	if sort != "" {
		// Handle descending sort (prefix with -)
		if strings.HasPrefix(sort, "-") {
			sortField := strings.TrimPrefix(sort, "-")
			// Map sort fields to database columns
			sortField = r.mapSortField(sortField)
			query = query.Order(sortField + " DESC")
		} else {
			sortField := r.mapSortField(sort)
			query = query.Order(sortField + " ASC")
		}
	} else {
		query = query.Order("satpen.created_at DESC")
	}

	err := query.Find(&satpen).Error
	return satpen, total, err
}

func (r *satpenRepository) FindByID(id uint) (*models.Satpen, error) {
	var satpen models.Satpen
	err := r.db.Preload("Provinsi").
		Preload("Kabupaten").
		Preload("Jenjang").
		Preload("Kategori").
		Preload("PengurusCabang").
		Preload("PDPTK", func(db *gorm.DB) *gorm.DB {
			return db.Order("tapel DESC").Limit(1)
		}).
		First(&satpen, id).Error

	if err != nil {
		return nil, err
	}
	return &satpen, nil
}

func (r *satpenRepository) FindByNPSN(npsn string) (*models.Satpen, error) {
	var satpen models.Satpen
	err := r.db.Preload("Provinsi").
		Preload("Kabupaten").
		Preload("Jenjang").
		Preload("Kategori").
		Preload("PengurusCabang").
		Preload("PDPTK", func(db *gorm.DB) *gorm.DB {
			return db.Order("tapel DESC").Limit(1)
		}).
		Where("npsn = ?", npsn).
		First(&satpen).Error

	if err != nil {
		return nil, err
	}
	return &satpen, nil
}

func (r *satpenRepository) GetStatistics(filters map[string]interface{}) (*models.SatpenStatistics, error) {
	stats := &models.SatpenStatistics{}

	query := r.db.Model(&models.Satpen{})
	query = r.applyFilters(query, filters)

	// Total satpen
	if err := query.Count(&stats.TotalSatpen).Error; err != nil {
		return nil, err
	}

	// Total siswa and guru from PDPTK (latest data)
	var sums struct {
		TotalSiswa int64
		TotalGuru  int64
	}

	subQuery := r.db.Table("pdptk").
		Select("id_satpen, MAX(tapel) as max_tapel").
		Group("id_satpen")

	if err := r.db.Table("pdptk").
		Select("COALESCE(SUM(jml_pd), 0) as total_siswa, COALESCE(SUM(jml_guru), 0) as total_guru").
		Joins("INNER JOIN (?) as latest ON pdptk.id_satpen = latest.id_satpen AND pdptk.tapel = latest.max_tapel", subQuery).
		Scan(&sums).Error; err != nil {
		return nil, err
	}
	stats.TotalSiswa = sums.TotalSiswa
	stats.TotalGuru = sums.TotalGuru

	// Total provinsi
	if err := r.db.Model(&models.Provinsi{}).Count(&stats.TotalProvinsi).Error; err != nil {
		return nil, err
	}

	// By Jenjang
	jenjangCounts, err := r.CountByJenjang(filters)
	if err != nil {
		return nil, err
	}
	stats.ByJenjang = make(map[string]models.JenjangStats)
	for _, jc := range jenjangCounts {
		stats.ByJenjang[jc.Jenjang] = models.JenjangStats{
			Count: jc.Count,
			Siswa: jc.Siswa,
			Guru:  jc.Guru,
		}
	}

	// By Akreditasi (based on kategori)
	akreditasiCounts, err := r.CountByAkreditasi(filters)
	if err != nil {
		return nil, err
	}
	stats.ByAkreditasi = make(map[string]int64)
	for _, ac := range akreditasiCounts {
		stats.ByAkreditasi[ac.Akreditasi] = ac.Count
	}

	// Top Provinsi
	topProvinsi, err := r.GetTopProvinsi(5)
	if err != nil {
		return nil, err
	}
	stats.TopProvinsi = topProvinsi

	return stats, nil
}

func (r *satpenRepository) CountByJenjang(filters map[string]interface{}) ([]models.JenjangCount, error) {
	var results []models.JenjangCount

	query := r.db.Table("satpen").
		Select("jenjang_pendidikan.nm_jenjang as jenjang, COUNT(*) as count, COALESCE(SUM(pdptk.jml_pd), 0) as siswa, COALESCE(SUM(pdptk.jml_guru), 0) as guru").
		Joins("INNER JOIN jenjang_pendidikan ON jenjang_pendidikan.id_jenjang = satpen.id_jenjang").
		Joins("LEFT JOIN (SELECT id_satpen, jml_pd, jml_guru FROM pdptk WHERE (id_satpen, tapel) IN (SELECT id_satpen, MAX(tapel) FROM pdptk GROUP BY id_satpen)) as pdptk ON pdptk.id_satpen = satpen.id_satpen").
		Group("jenjang_pendidikan.id_jenjang, jenjang_pendidikan.nm_jenjang")

	query = r.applyFilters(query, filters)

	err := query.Scan(&results).Error
	return results, err
}

func (r *satpenRepository) CountByAkreditasi(filters map[string]interface{}) ([]models.AkreditasiCount, error) {
	var results []models.AkreditasiCount

	query := r.db.Table("satpen").
		Select("COALESCE(kategori_satpen.nm_kategori, 'Belum Terakreditasi') as akreditasi, COUNT(*) as count").
		Joins("LEFT JOIN kategori_satpen ON kategori_satpen.id_kategori = satpen.id_kategori").
		Group("kategori_satpen.nm_kategori")

	query = r.applyFilters(query, filters)

	err := query.Scan(&results).Error
	return results, err
}

func (r *satpenRepository) GetTopProvinsi(limit int) ([]models.ProvinsiStats, error) {
	var results []models.ProvinsiStats

	err := r.db.Table("satpen").
		Select("provinsi.nm_prov as provinsi, COUNT(*) as count").
		Joins("INNER JOIN provinsi ON provinsi.id_prov = satpen.id_prov").
		Group("provinsi.id_prov, provinsi.nm_prov").
		Order("count DESC").
		Limit(limit).
		Scan(&results).Error

	return results, err
}

func (r *satpenRepository) applyFilters(query *gorm.DB, filters map[string]interface{}) *gorm.DB {
	// Filter by jenjang (by name)
	if jenjang, ok := filters["jenjang"].(string); ok && jenjang != "" {
		query = query.Where("satpen.id_jenjang IN (SELECT id_jenjang FROM jenjang_pendidikan WHERE nm_jenjang = ?)", jenjang)
	}

	// Filter by provinsi (by name or ID)
	if provinsi, ok := filters["provinsi"].(string); ok && provinsi != "" {
		query = query.Where("satpen.id_prov IN (SELECT id_prov FROM provinsi WHERE nm_prov LIKE ? OR id_prov = ?)", "%"+provinsi+"%", provinsi)
	}

	// Filter by kabupaten (by name or ID)
	if kabupaten, ok := filters["kabupaten"].(string); ok && kabupaten != "" {
		query = query.Where("satpen.id_kab IN (SELECT id_kab FROM kabupaten WHERE nama_kab LIKE ? OR id_kab = ?)", "%"+kabupaten+"%", kabupaten)
	}

	// Search by name or address
	if search, ok := filters["search"].(string); ok && search != "" {
		searchTerm := "%" + search + "%"
		query = query.Where("satpen.nm_satpen LIKE ? OR satpen.alamat LIKE ?", searchTerm, searchTerm)
	}

	// Filter by akreditasi (kategori)
	if akreditasi, ok := filters["akreditasi"].(string); ok && akreditasi != "" {
		query = query.Where("satpen.id_kategori IN (SELECT id_kategori FROM kategori_satpen WHERE nm_kategori = ?)", akreditasi)
	}

	// Filter by status (default: setujui = aktif)
	if status, ok := filters["status"].(string); ok && status != "" {
		if status == "aktif" {
			query = query.Where("satpen.status = ?", "setujui")
		} else if status == "non-aktif" {
			query = query.Where("satpen.status IN (?)", []string{"expired", "revisi", "permohonan"})
		} else {
			query = query.Where("satpen.status = ?", status)
		}
	} else {
		// Default: only show approved (aktif)
		query = query.Where("satpen.status = ?", "setujui")
	}

	// Filter by verified
	if verified, ok := filters["verified"].(bool); ok {
		if verified {
			query = query.Where("satpen.status = ?", "setujui")
		} else {
			query = query.Where("satpen.status != ?", "setujui")
		}
	}

	return query
}

func (r *satpenRepository) mapSortField(field string) string {
	// Map API sort fields to database columns
	mapping := map[string]string{
		"nama":          "satpen.nm_satpen",
		"jumlah_siswa":  "pdptk.jml_pd",
		"tahun_berdiri": "satpen.thn_berdiri",
		"created_at":    "satpen.created_at",
		"updated_at":    "satpen.updated_at",
	}

	if mapped, ok := mapping[field]; ok {
		return mapped
	}
	return "satpen." + field
}
