package service

import (
	"satpen-api/internal/models"
	"satpen-api/internal/repository"
)

type MasterService interface {
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

type masterService struct {
	repo repository.MasterRepository
}

func NewMasterService(repo repository.MasterRepository) MasterService {
	return &masterService{repo: repo}
}

// Provinsi Methods
func (s *masterService) GetAllProvinsi(search string) ([]models.Provinsi, error) {
	return s.repo.GetAllProvinsi(search)
}

func (s *masterService) GetProvinsiByID(id uint) (*models.Provinsi, error) {
	return s.repo.GetProvinsiByID(id)
}

// Kabupaten Methods
func (s *masterService) GetAllKabupaten(provinsiID uint, search string) ([]models.Kabupaten, error) {
	return s.repo.GetAllKabupaten(provinsiID, search)
}

func (s *masterService) GetKabupatenByID(id uint) (*models.Kabupaten, error) {
	return s.repo.GetKabupatenByID(id)
}

// Pengurus Cabang Methods
func (s *masterService) GetAllPengurusCabang(filters map[string]interface{}, page, limit int) ([]models.PengurusCabang, int64, error) {
	return s.repo.GetAllPengurusCabang(filters, page, limit)
}

func (s *masterService) GetPengurusCabangByID(id uint) (*models.PengurusCabang, error) {
	return s.repo.GetPengurusCabangByID(id)
}

// Jenjang Pendidikan Methods
func (s *masterService) GetAllJenjangPendidikan(search string) ([]models.JenjangPendidikan, error) {
	return s.repo.GetAllJenjangPendidikan(search)
}

func (s *masterService) GetJenjangPendidikanByID(id uint) (*models.JenjangPendidikan, error) {
	return s.repo.GetJenjangPendidikanByID(id)
}
