package models

import "time"

type Kabupaten struct {
	IDKab     uint      `json:"id" gorm:"column:id_kab;primaryKey"`
	IDProv    uint      `json:"id_prov" gorm:"column:id_prov;not null"`
	Provinsi  *Provinsi `json:"provinsi,omitempty" gorm:"foreignKey:IDProv;references:IDProv"`
	NamaKab   string    `json:"nama" gorm:"column:nama_kab;size:255;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (Kabupaten) TableName() string {
	return "kabupaten"
}
