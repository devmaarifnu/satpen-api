package repository

import (
	"satpen-api/internal/models"

	"gorm.io/gorm"
)

type MasterRepository interface {
	// Provinsi
	GetAllProvinsi(search string) ([]models.Provinsi, error)
	GetProvinsiByID(id uint) (*models.Provinsi, error)

	// Kabupaten
	GetAllKabupaten(provinsiID uint, search string) ([]models.Kabupaten, error)
	GetKabupatenByID(id uint) (*models.Kabupaten, error)

	// Pengurus Cabang
	GetAllPengurusCabang(filters map[string]interface{}, page, limit int) ([]models.PengurusCabang, int64, error)
	GetPengurusCabangByID(id uint) (*models.PengurusCabang, error)

	// Jenjang Pendidikan
	GetAllJenjangPendidikan(search string) ([]models.JenjangPendidikan, error)
	GetJenjangPendidikanByID(id uint) (*models.JenjangPendidikan, error)
}

type masterRepository struct {
	db *gorm.DB
}

func NewMasterRepository(db *gorm.DB) MasterRepository {
	return &masterRepository{db: db}
}

// Provinsi Methods
func (r *masterRepository) GetAllProvinsi(search string) ([]models.Provinsi, error) {
	var provinsi []models.Provinsi
	query := r.db.Model(&models.Provinsi{})

	if search != "" {
		query = query.Where("nm_prov LIKE ?", "%"+search+"%")
	}

	query = query.Order("nm_prov ASC")
	err := query.Find(&provinsi).Error
	return provinsi, err
}

func (r *masterRepository) GetProvinsiByID(id uint) (*models.Provinsi, error) {
	var provinsi models.Provinsi
	err := r.db.First(&provinsi, id).Error
	return &provinsi, err
}

// Kabupaten Methods
func (r *masterRepository) GetAllKabupaten(provinsiID uint, search string) ([]models.Kabupaten, error) {
	var kabupaten []models.Kabupaten
	query := r.db.Model(&models.Kabupaten{}).Preload("Provinsi")

	if provinsiID > 0 {
		query = query.Where("id_prov = ?", provinsiID)
	}

	if search != "" {
		query = query.Where("nama_kab LIKE ?", "%"+search+"%")
	}

	query = query.Order("nama_kab ASC")
	err := query.Find(&kabupaten).Error
	return kabupaten, err
}

func (r *masterRepository) GetKabupatenByID(id uint) (*models.Kabupaten, error) {
	var kabupaten models.Kabupaten
	err := r.db.Preload("Provinsi").First(&kabupaten, id).Error
	return &kabupaten, err
}

// Pengurus Cabang Methods
func (r *masterRepository) GetAllPengurusCabang(filters map[string]interface{}, page, limit int) ([]models.PengurusCabang, int64, error) {
	var pengurusCabang []models.PengurusCabang
	var total int64

	query := r.db.Model(&models.PengurusCabang{}).Preload("Provinsi")

	// Apply filters
	if provinsiID, ok := filters["provinsi_id"].(uint); ok && provinsiID > 0 {
		query = query.Where("id_prov = ?", provinsiID)
	}

	if search, ok := filters["search"].(string); ok && search != "" {
		query = query.Where("nama_pc LIKE ?", "%"+search+"%")
	}

	// Count total
	query.Count(&total)

	// Pagination
	offset := (page - 1) * limit
	query = query.Offset(offset).Limit(limit)

	// Order
	query = query.Order("nama_pc ASC")

	err := query.Find(&pengurusCabang).Error
	return pengurusCabang, total, err
}

func (r *masterRepository) GetPengurusCabangByID(id uint) (*models.PengurusCabang, error) {
	var pengurusCabang models.PengurusCabang
	err := r.db.Preload("Provinsi").First(&pengurusCabang, id).Error
	return &pengurusCabang, err
}

// Jenjang Pendidikan Methods
func (r *masterRepository) GetAllJenjangPendidikan(search string) ([]models.JenjangPendidikan, error) {
	var jenjang []models.JenjangPendidikan
	query := r.db.Model(&models.JenjangPendidikan{})

	if search != "" {
		query = query.Where("nm_jenjang LIKE ?", "%"+search+"%")
	}

	query = query.Order("nm_jenjang ASC")
	err := query.Find(&jenjang).Error
	return jenjang, err
}

func (r *masterRepository) GetJenjangPendidikanByID(id uint) (*models.JenjangPendidikan, error) {
	var jenjang models.JenjangPendidikan
	err := r.db.First(&jenjang, id).Error
	return &jenjang, err
}
