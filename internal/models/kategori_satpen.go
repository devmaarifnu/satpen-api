package models

import "time"

type KategoriSatpen struct {
	IDKategori uint      `json:"id" gorm:"column:id_kategori;primaryKey"`
	NmKategori string    `json:"nama" gorm:"column:nm_kategori;type:enum('A','B','C','D');not null"`
	Konotasi   string    `json:"konotasi" gorm:"column:konotasi;size:100;not null"`
	Keterangan string    `json:"keterangan,omitempty" gorm:"column:keterangan;size:255"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (KategoriSatpen) TableName() string {
	return "kategori_satpen"
}
