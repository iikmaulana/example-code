package controller

import (
	"github.com/iikmaulana/example-code/composite/models"
	"github.com/iikmaulana/example-code/composite/service"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
)

type kelasUsecase struct {
	kelasRepo service.KelasRepo
}

func NewKelasUsecase(kelasRepo service.KelasRepo) service.KelasUsecase {
	return kelasUsecase{kelasRepo: kelasRepo}
}

func (r kelasUsecase) GetKelasByIdUsecase(id string) (result models.KelasResult, serr serror.SError) {

	//todo memanggil repository kelas
	tmpResult, err := r.kelasRepo.GetKelasByIdRepo(id)
	if err != nil {
		return result, err
	}

	result = tmpResult
	return result, nil
}

func (r kelasUsecase) CreateKelasUsecase(form models.CreateKelasRequest) (result string, serr serror.SError) {

	//todo validasi request saat add kelas
	if form.NamaKelas == "" {
		serr := serror.New("Nama kelas tidak boleh kosong")
		return result, serr
	}

	//todo memanggil repository kelas
	tmpResult, err := r.kelasRepo.CreateKelasRepo(form)
	if err != nil {
		return result, err
	}

	result = tmpResult
	return result, nil
}
