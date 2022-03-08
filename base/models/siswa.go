package models

type CreateSiswaRequest struct {
	Nama    string `json:"nama"`
	Alamat  string `json:"alamat"`
	KelasId string `json:"kelas_id"`
}

type SiswaResult struct {
	Id      string `db:"id" json:"id"`
	Nama    string `db:"nama" json:"nama"`
	Alamat  string `db:"alamat" json:"alamat"`
	KelasId string `db:"kelas_id" json:"kelas_id"`
}

type ListSiswaResult struct {
	TotalData        int64         `json:"total_data"`
	TotalDataPerpage int           `json:"total_data_perpage"`
	From             int           `json:"from"`
	To               int           `json:"to"`
	TotalPage        int           `json:"total_page"`
	Data             []SiswaResult `json:"data"`
}
