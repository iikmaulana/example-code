package impl

import (
	"github.com/google/uuid"
	"github.com/iikmaulana/example-code/base/models"
	"github.com/iikmaulana/example-code/base/service"
	"github.com/iikmaulana/example-code/base/service/repository/query"
	"github.com/jmoiron/sqlx"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
)

type userRepo struct {
	DB *sqlx.DB
}

func NewKelasRepository(db *sqlx.DB) service.KelasRepo {
	return userRepo{DB: db}
}

func (r userRepo) GetKelasByIdRepo(id string) (result models.KelasResult, serr serror.SError) {
	rows, err := r.DB.Queryx(query.GetKelasById, id)
	if err != nil {
		return result, serror.New(err.Error())
	}

	defer rows.Close()
	for rows.Next() {
		if err := rows.StructScan(&result); err != nil {
			return result, serror.New("Failed scan for struct models")
		}
	}

	return result, nil
}

func (r userRepo) CreateKelasRepo(form models.CreateKelasRequest) (result string, serr serror.SError) {
	err := r.DB.QueryRow(query.CreateKelas, uuid.New().String(), form.NamaKelas).Scan(&result)
	if err != nil {
		return result, serror.New(err.Error())
	}

	return result, nil
}
