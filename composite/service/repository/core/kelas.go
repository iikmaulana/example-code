package core

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/iikmaulana/example-code/base/packets"
	"github.com/iikmaulana/example-code/composite/models"
	"github.com/iikmaulana/example-code/composite/service"
	"github.com/jmoiron/sqlx"
	"github.com/uzzeet/uzzeet-gateway/controller"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/logger"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
	"time"
)

type kelasRepository struct {
	client packets.KelasClient
	DB     *sqlx.DB
}

func NewKelasRepository(reg *controller.Registry, db *sqlx.DB) (res service.KelasRepo, serr serror.SError) {
	obj := kelasRepository{}
	i := 3

	//todo melakukan perulangan untuk memanggil registry dari redis
	for {
		if i <= 0 {
			serr = serror.Newc("Cannot connect to core Kelas",
				"[repository][core][kelas] while dialing core Kelas")
			break
		}

		//todo memanggil registry dari redis berdasarkan app key
		conn, err := reg.GetConnection("base_examaple_code")
		if err != nil {
			logger.Warn("[repository][core][kelas] while dial core Kelas")
			time.Sleep(1 * time.Second)

			i--
			continue
		}
		obj.client = packets.NewKelasClient(conn)
		break
	}

	obj.DB = db
	return obj, serr
}

func (r kelasRepository) GetKelasByIdRepo(id string) (result models.KelasResult, serr serror.SError) {

	//todo memanggil RPC get kelas berdasarkan ID
	output, err := r.client.GetKelasByIdUsecase(context.Background(), &packets.KelasRequestByID{
		KelasID: id,
	})

	//todo pengecekan error saat memanggil RPC kelas
	if err != nil {
		serrFmt := fmt.Sprintf("[service][repository][core][kelas] while grpc GetKelasByIdUsecase: %+v", err)
		logger.Err(serrFmt)
		return result, serror.NewFromErrorc(err, serrFmt)
	}

	//todo pengecekan return hasil RPC dan merubah ke dalam struct model
	if output.GetStatus() == 1 {
		err := json.Unmarshal(output.GetData().Value, &result)
		if err != nil {
			return result, serror.NewFromError(err)
		}
	}

	return result, nil
}

func (r kelasRepository) CreateKelasRepo(form models.CreateKelasRequest) (result string, serr serror.SError) {

	//todo merubah request model ke dalam bentuk byte dilakukan karena RPC menggunakan byte
	tmpByte, err := json.Marshal(form)
	if err != nil {
		return result, serror.NewFromError(err)
	}

	//todo membentuk request model yang diterima oleh RPC
	data := any.Any{
		Value: tmpByte,
	}

	//todo memanggil RPC add kelas
	output, err := r.client.CreateKelasUsecase(context.Background(), &packets.KelasRequest{
		Data: &data,
	})

	//todo pengecekan error saat memanggil RPC kelas
	if err != nil {
		serrFmt := fmt.Sprintf("[service][repository][core][kelas] while grpc CreateKelasUsecase: %+v", err)
		logger.Err(serrFmt)
		return result, serror.NewFromErrorc(err, serrFmt)
	}

	//todo pengecekan return hasil RPC dan merubah ke dalam struct model
	if output.GetStatus() == 1 {
		result = string(output.GetData().Value)
	}

	return result, nil
}
