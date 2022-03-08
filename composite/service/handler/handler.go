package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/iikmaulana/example-code/composite/service"
)

type gatewayHandler struct {
	service      *gin.Engine
	kelasUsecase service.KelasUsecase
	siswaUsecase service.SiswaUsecase
}

func NewGatewayHandler(svc *gin.Engine,
	kelasUsecase service.KelasUsecase,
	siswaUsecase service.SiswaUsecase,
) {
	h := gatewayHandler{
		service:      svc,
		kelasUsecase: kelasUsecase,
		siswaUsecase: siswaUsecase,
	}

	h.initUmUsage()

}
