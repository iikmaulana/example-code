package impl

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/iikmaulana/example-code/base/lib"
	"github.com/iikmaulana/example-code/base/models"
	"github.com/iikmaulana/example-code/base/service"
	"github.com/iikmaulana/example-code/base/service/repository/query"
	"github.com/jmoiron/sqlx"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
)

type siswaRepo struct {
	DB *sqlx.DB
}

func NewSiswaRepository(db *sqlx.DB) service.SiswaRepo {
	return siswaRepo{DB: db}
}

func (r siswaRepo) GetListSiswaRepo(form models.FilterParams) (result []models.SiswaResult, metas models.FormMetaData, serr serror.SError) {
	var totalData int64
	var offset int
	var paginate models.FormMetaData

	tmpQuery := query.GetListSiswa

	if form.Search != "" {
		if len(form.Filter) > 0 {
			for _, v := range form.Filter {
				tmpFilter := " and lower(replace(" + v + ", ' ', '')) like '%' || lower(replace('" + form.Search + "', ' ', '')) || '%'"
				tmpQuery = tmpQuery + tmpFilter
			}
		} else {
			return result, metas, serror.New("Tentukan filter yang akan di cari")
		}
	}

	queryTotalData := fmt.Sprintf(`with x as (%v) select count(*) as total_data from x`, tmpQuery)
	if err := r.DB.QueryRow(queryTotalData).Scan(&totalData); err != nil {
		return result, metas, serror.NewFromError(err)
	}

	paginate, offset = lib.PaginateFromQuery(totalData, int(form.Page), int(form.Limit))
	tmpQuery = tmpQuery + fmt.Sprintf(" limit %v offset %v ", form.Limit, offset)

	fmt.Println(tmpQuery)
	rows, err := r.DB.Queryx(tmpQuery)
	if err != nil {
		return result, metas, serror.New(err.Error())
	}

	defer rows.Close()
	for rows.Next() {
		tmpResult := models.SiswaResult{}
		if err := rows.StructScan(&tmpResult); err != nil {
			return result, metas, serror.New("Failed scan for struct models")
		}
		result = append(result, tmpResult)
	}

	metas = paginate
	metas.TotalDataPerpage = len(result)
	return result, metas, nil
}

func (r siswaRepo) CreateSiswaRepo(form models.CreateSiswaRequest) (result string, serr serror.SError) {
	err := r.DB.QueryRow(query.CreateSiswa, uuid.New().String(), form.Nama, form.Alamat, form.KelasId).Scan(&result)
	if err != nil {
		return result, serror.New(err.Error())
	}

	return result, nil
}
