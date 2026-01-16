package models

import "time"

type PDPTK struct {
	ID            int        `json:"id" gorm:"column:id;primaryKey"`
	IDSatpen      *uint      `json:"id_satpen,omitempty" gorm:"column:id_satpen"`
	Tapel         string     `json:"tapel,omitempty" gorm:"column:tapel;size:10"`
	PDLK          int        `json:"pd_lk" gorm:"column:pd_lk;default:0"`
	PDPR          int        `json:"pd_pr" gorm:"column:pd_pr;default:0"`
	JmlPD         int        `json:"jumlah_siswa" gorm:"column:jml_pd;default:0"`
	GuruLK        int        `json:"guru_lk" gorm:"column:guru_lk;default:0"`
	GuruPR        int        `json:"guru_pr" gorm:"column:guru_pr;default:0"`
	JmlGuru       int        `json:"jumlah_guru" gorm:"column:jml_guru;default:0"`
	TendikLK      int        `json:"tendik_lk" gorm:"column:tendik_lk;default:0"`
	TendikPR      int        `json:"tendik_pr" gorm:"column:tendik_pr;default:0"`
	JmlTendik     int        `json:"jumlah_tendik" gorm:"column:jml_tendik;default:0"`
	LastSinkron   *time.Time `json:"last_sinkron,omitempty" gorm:"column:last_sinkron"`
	StatusSinkron int        `json:"status_sinkron" gorm:"column:status_sinkron;type:tinyint"`
}

func (PDPTK) TableName() string {
	return "pdptk"
}
