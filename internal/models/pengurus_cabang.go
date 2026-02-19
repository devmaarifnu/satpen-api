package models

import "time"

type PengurusCabang struct {
	IDPC      uint      `json:"id" gorm:"column:id_pc;primaryKey"`
	IDProv    uint      `json:"id_prov" gorm:"column:id_prov;not null"`
	Provinsi  *Provinsi `json:"provinsi,omitempty" gorm:"foreignKey:IDProv;references:IDProv"`
	KodeKab   string    `json:"kode_kab" gorm:"column:kode_kab;size:10;not null"`
	NamaPC    string    `json:"nama" gorm:"column:nama_pc;size:255;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (PengurusCabang) TableName() string {
	return "pengurus_cabang"
}
