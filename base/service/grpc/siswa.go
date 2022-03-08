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

type SiswaHandler struct {
	siswaUsecase service.SiswaUsecase
}

func NewSiswaHandler(siswaUsecase service.SiswaUsecase) SiswaHandler {
	return SiswaHandler{
		siswaUsecase: siswaUsecase}
}

func (s SiswaHandler) GetListSiswaUsecase(ctx context.Context, in *packets.SiswaRequest) (output *packets.OutputSiswa, serr error) {
	output = &packets.OutputSiswa{
		Status: 0,
	}

	tmpForm := models.FilterParams{}
	errx := json.Unmarshal(in.GetData().Value, &tmpForm)
	if errx != nil {
		return output, serror.NewFromError(errx)
	}

	result, serr := s.siswaUsecase.GetListSiswaUsecase(tmpForm)
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

func (s SiswaHandler) CreateSiswaUsecase(ctx context.Context, in *packets.SiswaRequest) (output *packets.OutputSiswa, serr error) {
	output = &packets.OutputSiswa{
		Status: 0,
	}

	request := models.CreateSiswaRequest{}
	errx := json.Unmarshal(in.GetData().Value, &request)
	if errx != nil {
		return output, serror.NewFromError(errx)
	}

	id, serr := s.siswaUsecase.CreateSiswaUsecase(request)
	if serr != nil {
		return output, serr
	}

	output.Status = 1
	output.Data = &any.Any{
		Value: []byte(id),
	}

	return output, nil
}
