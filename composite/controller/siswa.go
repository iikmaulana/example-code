package controller

import (
	"github.com/iikmaulana/example-code/composite/models"
	"github.com/iikmaulana/example-code/composite/service"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
)

type siswaUsecase struct {
	siswaRepo service.SiswaRepo
	kelasRepo service.KelasRepo
}

func NewSiswaUsecase(siswaRepo service.SiswaRepo, kelasRepo service.KelasRepo) service.SiswaUsecase {
	return siswaUsecase{siswaRepo: siswaRepo, kelasRepo: kelasRepo}
}

func (r siswaUsecase) GetListSiswaUsecase(form models.FilterParams) (result models.ListSiswaResult, serr serror.SError) {

	//todo memanggil repository siswa
	tmpResult, err := r.siswaRepo.GetListSiswaRepo(form)
	if err != nil {
		return result, err
	}

	//todo melakukan mapping payload result
	tmpData := []models.SiswaResult{}
	for _, v := range tmpResult.Data {
		tmpKelas, _ := r.kelasRepo.GetKelasByIdRepo(v.KelasId)
		tmpDataObject := models.SiswaResult{
			Id:     v.Id,
			Nama:   v.Nama,
			Alamat: v.Alamat,
			Kelas:  *tmpKelas.NamaKelas,
		}

		tmpData = append(tmpData, tmpDataObject)
	}

	result.Data = tmpData
	result.TotalData = tmpResult.TotalData
	result.TotalDataPerpage = len(tmpData)
	result.From = tmpResult.From
	result.To = tmpResult.To
	result.TotalPage = tmpResult.TotalPage
	return result, nil
}

func (r siswaUsecase) CreateSiswaUsecase(form models.CreateSiswaRequest) (result string, serr serror.SError) {

	//todo melakukan pengecekan id kelas berdasarkan kelas_id dari RPC kelas
	if form.KelasId != "" {
		_, err := r.kelasRepo.GetKelasByIdRepo(form.KelasId)
		if err != nil {
			return result, err
		}
	}

	//todo validasi request saat add siswa
	if form.Nama == "" {
		serr := serror.New("Nama siswa tidak boleh kosong")
		return result, serr
	}

	//todo memanggil repository siswa
	tmpResult, err := r.siswaRepo.CreateSiswaRepo(form)
	if err != nil {
		return result, err
	}

	result = tmpResult
	return result, nil
}
