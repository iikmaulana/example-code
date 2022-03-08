package models

type CreateKelasRequest struct {
	NamaKelas string `json:"nama_kelas"`
}

type KelasResult struct {
	ID        string  `db:"id" json:"id"`
	NamaKelas *string `db:"username" json:"nama_kelas"`
}

type ListKelasResult struct {
	TotalData        int64         `json:"total_data"`
	TotalDataPerpage int           `json:"total_data_perpage"`
	From             int           `json:"from"`
	To               int           `json:"to"`
	TotalPage        int           `json:"total_page"`
	Data             []KelasResult `json:"data"`
}
