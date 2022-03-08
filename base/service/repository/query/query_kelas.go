package query

const (
	CreateKelas  = `INSERT INTO example_code.public.kelas (id, nama_kelas) VALUES ($1, $2) returning id`
	GetKelasById = `SELECT id, nama_kelas FROM example_code.public.kelas where id = $1`
)
