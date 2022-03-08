package service

import (
	"github.com/iikmaulana/example-code/base/models"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
)

type KelasRepo interface {
	GetKelasByIdRepo(id string) (result models.KelasResult, serr serror.SError)
	CreateKelasRepo(form models.CreateKelasRequest) (result string, serr serror.SError)
}

type SiswaRepo interface {
	GetListSiswaRepo(form models.FilterParams) (result []models.SiswaResult, metas models.FormMetaData, serr serror.SError)
	CreateSiswaRepo(form models.CreateSiswaRequest) (result string, serr serror.SError)
}
