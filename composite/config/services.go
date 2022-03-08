package config

import (
	"github.com/iikmaulana/example-code/composite/controller"
	"github.com/iikmaulana/example-code/composite/service/handler"
	"github.com/iikmaulana/example-code/composite/service/repository/core"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
)

func (cfg Config) InitService() serror.SError {

	//todo memanggil connection rpc kelas
	kelasRepo, serr := core.NewKelasRepository(cfg.Registry, cfg.DB)
	if serr != nil {
		return serr
	}

	//todo memanggil connection rpc siswa
	siswaRepo, serr := core.NewSiswaRepository(cfg.Registry, cfg.DB)
	if serr != nil {
		return serr
	}

	//todo memasukan rpc kedalam controller
	kelasUsecase := controller.NewKelasUsecase(kelasRepo)
	siswaUsecase := controller.NewSiswaUsecase(siswaRepo, kelasRepo)

	//todo memasukan controller kedalam handler
	handler.NewGatewayHandler(cfg.Gin, kelasUsecase, siswaUsecase)

	return nil
}
