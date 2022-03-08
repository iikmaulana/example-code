package config

import (
	"github.com/iikmaulana/example-code/base/controller"
	"github.com/iikmaulana/example-code/base/packets"
	"github.com/iikmaulana/example-code/base/service/grpc"
	"github.com/iikmaulana/example-code/base/service/repository/impl"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
)

func (cfg Config) InitService() serror.SError {

	kelasRepo := impl.NewKelasRepository(cfg.DB)
	kelasUsecase := controller.NewKelasUsecase(kelasRepo)

	siswaRepo := impl.NewSiswaRepository(cfg.DB)
	siswaUsecase := controller.NewSiswaUsecase(siswaRepo)

	packets.RegisterKelasServer(cfg.Server.Instance(), grpc.NewKelasHandler(kelasUsecase))
	packets.RegisterSiswaServer(cfg.Server.Instance(), grpc.NewSiswaHandler(siswaUsecase))
	return nil
}
