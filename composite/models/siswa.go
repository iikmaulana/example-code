package models

type CreateSiswaRequest struct {
	Nama    string `json:"nama"`
	Alamat  string `json:"alamat"`
	KelasId string `json:"kelas_id"`
}

type SiswaResultRpc struct {
	Id      string `json:"id"`
	Nama    string `json:"nama"`
	Alamat  string `json:"alamat"`
	KelasId string `json:"kelas_id"`
}

type ListSiswaResultRpc struct {
	TotalData        int64            `json:"total_data"`
	TotalDataPerpage int              `json:"total_data_perpage"`
	From             int              `json:"from"`
	To               int              `json:"to"`
	TotalPage        int              `json:"total_page"`
	Data             []SiswaResultRpc `json:"data"`
}

type SiswaResult struct {
	Id     string `json:"id"`
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
	Kelas  string `json:"kelas"`
}

type ListSiswaResult struct {
	TotalData        int64         `json:"total_data"`
	TotalDataPerpage int           `json:"total_data_perpage"`
	From             int           `json:"from"`
	To               int           `json:"to"`
	TotalPage        int           `json:"total_page"`
	Data             []SiswaResult `json:"data"`
}
