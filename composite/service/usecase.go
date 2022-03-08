package service

import (
	"github.com/iikmaulana/example-code/composite/models"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
)

type KelasUsecase interface {
	GetKelasByIdUsecase(id string) (result models.KelasResult, serr serror.SError)
	CreateKelasUsecase(form models.CreateKelasRequest) (result string, serr serror.SError)
}

type SiswaUsecase interface {
	GetListSiswaUsecase(form models.FilterParams) (result models.ListSiswaResult, serr serror.SError)
	CreateSiswaUsecase(form models.CreateSiswaRequest) (result string, serr serror.SError)
}
