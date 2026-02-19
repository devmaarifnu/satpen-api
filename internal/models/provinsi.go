package models

import "time"

type Provinsi struct {
	IDProv    uint      `json:"id" gorm:"column:id_prov;primaryKey"`
	Map       string    `json:"map,omitempty" gorm:"column:map;size:10"`
	KodeProv  string    `json:"kode" gorm:"column:kode_prov;size:10;not null"`
	NmProv    string    `json:"nama" gorm:"column:nm_prov;size:100;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (Provinsi) TableName() string {
	return "provinsi"
}
