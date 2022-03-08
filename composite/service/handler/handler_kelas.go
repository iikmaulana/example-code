package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/iikmaulana/example-code/composite/lib"
	"github.com/iikmaulana/example-code/composite/models"
	"net/http"
)

var svcidKelas = "kelas"

func (ox gatewayHandler) CreateKelas(ctx *gin.Context) {

	//todo jika ada validasi token
	/*token, err_t := lib.ClaimToken(ctx.Request.Header["Authorization"])
	if err_t != nil {
		serr := serror.NewFromError(err_t)
		lib.Response(http.StatusForbidden, serr.Error(), "", svcidKelas, "add_kelas", "add", "", ctx)
		return
	}

	if !lib.CheckArrayString(lib.AddUser, token.UserAccess) {
		serr := serror.New("Token access denied")
		lib.Response(http.StatusForbidden, serr.Error(), "", svcidKelas, "add_kelas", "add", "", ctx)
		return
	}*/

	//todo membuat model request add kelas
	tmpForm := models.CreateKelasRequest{
		NamaKelas: ctx.PostForm("nama_kelas"),
	}

	//todo memanggil controller saat add kelas
	result, err := ox.kelasUsecase.CreateKelasUsecase(tmpForm)
	if err != nil {
		lib.Response(http.StatusInternalServerError, err.Error(), "", svcidKelas, "add_kelas", "add", "", ctx)
		return
	}

	lib.Response(http.StatusOK, "", "", svcidKelas, "add_kelas", "add", result, ctx)
	return
}
