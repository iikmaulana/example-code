package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/iikmaulana/example-code/composite/lib"
	"github.com/iikmaulana/example-code/composite/models"
	"github.com/uzzeet/uzzeet-gateway/libs/helper"
	"net/http"
)

var svcidSiswa = "siswa"

func (ox gatewayHandler) CreateSiswa(ctx *gin.Context) {

	//todo jika ada validasi token
	/*token, err_t := lib.ClaimToken(ctx.Request.Header["Authorization"])
	if err_t != nil {
		serr := serror.NewFromError(err_t)
		lib.Response(http.StatusForbidden, serr.Error(), "", svcidSiswa, "add_kelas", "add", "", ctx)
		return
	}

	if !lib.CheckArrayString(lib.AddUser, token.UserAccess) {
		serr := serror.New("Token access denied")
		lib.Response(http.StatusForbidden, serr.Error(), "", svcidSiswa, "add_kelas", "add", "", ctx)
		return
	}*/

	//todo membuat model request
	tmpForm := models.CreateSiswaRequest{
		Nama:    ctx.PostForm("nama"),
		Alamat:  ctx.PostForm("alamat"),
		KelasId: ctx.PostForm("kelas_id"),
	}

	//todo memanggil conttroler add siswa
	result, err := ox.siswaUsecase.CreateSiswaUsecase(tmpForm)
	if err != nil {
		lib.Response(http.StatusInternalServerError, err.Error(), "", svcidSiswa, "add_kelas", "add", "", ctx)
		return
	}

	lib.Response(http.StatusOK, "", "", svcidSiswa, "add_kelas", "add", result, ctx)
	return
}

func (ox gatewayHandler) ListSiswa(ctx *gin.Context) {

	//todo jika ada validasi token
	/*token, err_t := lib.ClaimToken(ctx.Request.Header["Authorization"])
	if err_t != nil {
		serr := serror.NewFromError(err_t)
		lib.Response(http.StatusForbidden, serr.Error(), "", svcidSiswa, "add_kelas", "add", "", ctx)
		return
	}

	if !lib.CheckArrayString(lib.AddUser, token.UserAccess) {
		serr := serror.New("Token access denied")
		lib.Response(http.StatusForbidden, serr.Error(), "", svcidSiswa, "add_kelas", "add", "", ctx)
		return
	}*/

	//todo membuat model request
	tmpForm := models.FilterParams{
		Page:  helper.StringToInt(ctx.Query("page"), 1),
		Limit: helper.StringToInt(ctx.Query("limit"), 10),
	}

	//todo memanggil conttroler list siswa
	result, err := ox.siswaUsecase.GetListSiswaUsecase(tmpForm)
	if err != nil {
		lib.Response(http.StatusInternalServerError, err.Error(), "", svcidSiswa, "edit_user", "edit", "", ctx)
		return
	}

	lib.Response(http.StatusOK, "", "", svcidSiswa, "edit_user", "edit", result, ctx)
	return
}
