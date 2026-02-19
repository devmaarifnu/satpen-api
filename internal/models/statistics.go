package models

type SatpenStatistics struct {
	TotalSatpen   int64                     `json:"total_satpen"`
	TotalProvinsi int64                     `json:"total_provinsi"`
	TotalSiswa    int64                     `json:"total_siswa"`
	TotalGuru     int64                     `json:"total_guru"`
	ByJenjang     map[string]JenjangStats   `json:"by_jenjang"`
	ByAkreditasi  map[string]int64          `json:"by_akreditasi"`
	TopProvinsi   []ProvinsiStats           `json:"top_provinsi"`
}

type JenjangStats struct {
	Count int64 `json:"count"`
	Siswa int64 `json:"siswa"`
	Guru  int64 `json:"guru"`
}

type ProvinsiStats struct {
	Provinsi string `json:"provinsi"`
	Count    int64  `json:"count"`
}

type AkreditasiCount struct {
	Akreditasi string
	Count      int64
}

type JenjangCount struct {
	Jenjang string
	Count   int64
	Siswa   int64
	Guru    int64
}
