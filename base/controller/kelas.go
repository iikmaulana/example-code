package controller

import (
	"github.com/iikmaulana/example-code/base/models"
	"github.com/iikmaulana/example-code/base/service"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
)

type kelasUsecase struct {
	kelasRepo service.KelasRepo
}

func NewKelasUsecase(kelasRepo service.KelasRepo) service.KelasUsecase {
	return kelasUsecase{kelasRepo: kelasRepo}
}

func (f kelasUsecase) GetKelasByIdUsecase(id string) (result models.KelasResult, serr serror.SError) {
	result, err := f.kelasRepo.GetKelasByIdRepo(id)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (f kelasUsecase) CreateKelasUsecase(form models.CreateKelasRequest) (result string, serr serror.SError) {
	result, err := f.kelasRepo.CreateKelasRepo(form)
	if err != nil {
		return result, err
	}
	return result, nil
}
