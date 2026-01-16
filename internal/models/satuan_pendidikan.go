package models

import (
	"time"
)

type Satpen struct {
	IDSatpen       uint               `json:"id" gorm:"column:id_satpen;primaryKey"`
	IDUser         uint               `json:"-" gorm:"column:id_user;not null"`
	IDProv         uint               `json:"-" gorm:"column:id_prov;not null"`
	Provinsi       *Provinsi          `json:"provinsi,omitempty" gorm:"foreignKey:IDProv;references:IDProv"`
	IDKab          uint               `json:"-" gorm:"column:id_kab;not null"`
	Kabupaten      *Kabupaten         `json:"kabupaten,omitempty" gorm:"foreignKey:IDKab;references:IDKab"`
	IDPC           uint               `json:"-" gorm:"column:id_pc;not null"`
	PengurusCabang *PengurusCabang    `json:"pengurus_cabang,omitempty" gorm:"foreignKey:IDPC;references:IDPC"`
	IDKategori     *uint              `json:"-" gorm:"column:id_kategori"`
	Kategori       *KategoriSatpen    `json:"kategori,omitempty" gorm:"foreignKey:IDKategori;references:IDKategori"`
	IDJenjang      uint               `json:"-" gorm:"column:id_jenjang;not null"`
	Jenjang        *JenjangPendidikan `json:"jenjang,omitempty" gorm:"foreignKey:IDJenjang;references:IDJenjang"`
	NPSN           string             `json:"npsn" gorm:"column:npsn;size:45;uniqueIndex;not null"`
	NoRegistrasi   string             `json:"no_registrasi" gorm:"column:no_registrasi;size:45;uniqueIndex;not null"`
	NoUrut         string             `json:"no_urut" gorm:"column:no_urut;size:10;uniqueIndex;not null"`
	NmSatpen       string             `json:"nama" gorm:"column:nm_satpen;size:255;not null"`
	Yayasan        string             `json:"yayasan" gorm:"column:yayasan;size:255;not null"`
	Kepsek         string             `json:"kepala_sekolah,omitempty" gorm:"column:kepsek;size:100"`
	Telpon         string             `json:"phone,omitempty" gorm:"column:telpon;size:15"`
	Fax            string             `json:"fax,omitempty" gorm:"column:fax;size:15"`
	Email          string             `json:"email,omitempty" gorm:"column:email;size:100"`
	ThnBerdiri     int                `json:"tahun_berdiri,omitempty" gorm:"column:thn_berdiri;type:year"`
	Kecamatan      string             `json:"kecamatan" gorm:"column:kecamatan;size:255;not null"`
	Kelurahan      string             `json:"kelurahan" gorm:"column:kelurahan;size:255;not null"`
	Alamat         string             `json:"alamat" gorm:"column:alamat;type:text;not null"`
	AsetTanah      string             `json:"aset_tanah,omitempty" gorm:"column:aset_tanah;size:45"`
	NmPemilik      string             `json:"nama_pemilik,omitempty" gorm:"column:nm_pemilik;size:100"`
	TglRegistrasi  time.Time          `json:"tanggal_registrasi" gorm:"column:tgl_registrasi;not null"`
	ActivedDate    *time.Time         `json:"actived_date,omitempty" gorm:"column:actived_date"`
	Status         string             `json:"status" gorm:"column:status;type:enum('permohonan','revisi','proses dokumen','setujui','expired','perpanjangan');not null"`
	CreatedAt      time.Time          `json:"created_at" gorm:"column:created_at"`
	UpdatedAt      time.Time          `json:"updated_at" gorm:"column:updated_at"`

	// Data PDPTK dari relasi (untuk jumlah siswa dan guru)
	PDPTK          *PDPTK             `json:"pdptk,omitempty" gorm:"foreignKey:IDSatpen;references:IDSatpen"`

	// Virtual fields untuk compatibility dengan API response
	JumlahSiswa    uint               `json:"jumlah_siswa" gorm:"-"`
	JumlahGuru     uint               `json:"jumlah_guru" gorm:"-"`
	JumlahRombel   uint               `json:"jumlah_rombel" gorm:"-"`
	Akreditasi     string             `json:"akreditasi,omitempty" gorm:"-"`
	IsVerified     bool               `json:"is_verified" gorm:"-"`
	VerifiedAt     *time.Time         `json:"verified_at,omitempty" gorm:"-"`
}

func (Satpen) TableName() string {
	return "satpen"
}

// AfterFind hook to populate jumlah siswa and guru from PDPTK
func (s *Satpen) AfterFind() error {
	if s.PDPTK != nil {
		s.JumlahSiswa = uint(s.PDPTK.JmlPD)
		s.JumlahGuru = uint(s.PDPTK.JmlGuru)
		s.JumlahRombel = 0 // Not available in current schema
	}

	// Map kategori to akreditasi
	if s.Kategori != nil {
		s.Akreditasi = s.Kategori.NmKategori
	}

	// Set verified based on status
	s.IsVerified = s.Status == "setujui"
	if s.IsVerified && s.ActivedDate != nil {
		s.VerifiedAt = s.ActivedDate
	}

	return nil
}
