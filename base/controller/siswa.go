package controller

import (
	"fmt"
	"github.com/iikmaulana/example-code/base/models"
	"github.com/iikmaulana/example-code/base/service"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
)

type siswaUsecase struct {
	siswaRepo service.SiswaRepo
}

func NewSiswaUsecase(siswaRepo service.SiswaRepo) service.SiswaUsecase {
	return siswaUsecase{siswaRepo: siswaRepo}
}

func (f siswaUsecase) GetListSiswaUsecase(form models.FilterParams) (result models.ListSiswaResult, serr serror.SError) {
	tmpData, metas, err := f.siswaRepo.GetListSiswaRepo(form)
	if err != nil {
		return result, err
	}

	result.Data = tmpData
	result.TotalData = metas.TotalData
	result.TotalDataPerpage = len(result.Data)
	result.From = metas.From
	result.To = metas.To
	result.TotalPage = metas.TotalPage

	fmt.Println(result)

	return result, nil
}

func (f siswaUsecase) CreateSiswaUsecase(form models.CreateSiswaRequest) (result string, serr serror.SError) {
	result, err := f.siswaRepo.CreateSiswaRepo(form)
	if err != nil {
		return result, err
	}
	return result, nil
}
