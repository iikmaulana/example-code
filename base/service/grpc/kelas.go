package grpc

import (
	"context"
	"encoding/json"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/iikmaulana/example-code/base/models"
	"github.com/iikmaulana/example-code/base/packets"
	"github.com/iikmaulana/example-code/base/service"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
)

type KelasHandler struct {
	kelasUsecase service.KelasUsecase
}

func NewKelasHandler(kelasUsecase service.KelasUsecase) KelasHandler {
	return KelasHandler{
		kelasUsecase: kelasUsecase}
}

func (handler KelasHandler) GetKelasByIdUsecase(ctx context.Context, in *packets.KelasRequestByID) (output *packets.OutputKelas, serr error) {
	output = &packets.OutputKelas{
		Status: 0,
	}

	result, serr := handler.kelasUsecase.GetKelasByIdUsecase(in.KelasID)
	if serr != nil {
		return output, serr
	}

	byte, err := json.Marshal(result)
	if err != nil {
		return output, serror.NewFromError(err)
	}

	output.Status = 1
	output.Data = &any.Any{
		Value: byte,
	}

	return output, nil
}

func (handler KelasHandler) CreateKelasUsecase(ctx context.Context, in *packets.KelasRequest) (output *packets.OutputKelas, serr error) {
	output = &packets.OutputKelas{
		Status: 0,
	}

	request := models.CreateKelasRequest{}
	errx := json.Unmarshal(in.GetData().Value, &request)
	if errx != nil {
		return output, serror.NewFromError(errx)
	}

	id, serr := handler.kelasUsecase.CreateKelasUsecase(request)
	if serr != nil {
		return output, serr
	}

	output.Status = 1
	output.Data = &any.Any{
		Value: []byte(id),
	}

	return output, nil
}
