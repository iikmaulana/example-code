package query

const (
	CreateSiswa  = `INSERT INTO example_code.public.siswa (id, nama, alamat, kelas_id) VALUES ($1, $2, $3, $4) returning id`
	GetListSiswa = `SELECT id, nama, alamat, kelas_id FROM example_code.public.siswa`
)
