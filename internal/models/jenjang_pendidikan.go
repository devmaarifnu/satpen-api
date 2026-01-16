package models

import "time"

type JenjangPendidikan struct {
	IDJenjang  uint      `json:"id" gorm:"column:id_jenjang;primaryKey"`
	NmJenjang  string    `json:"nama" gorm:"column:nm_jenjang;size:45;not null"`
	Keterangan string    `json:"keterangan,omitempty" gorm:"column:keterangan;size:255"`
	Lembaga    string    `json:"lembaga,omitempty" gorm:"column:lembaga;type:enum('MADRASAH','SEKOLAH')"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (JenjangPendidikan) TableName() string {
	return "jenjang_pendidikan"
}
