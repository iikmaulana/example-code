package handler

func (ox gatewayHandler) initUmUsage() {

	router := ox.service.Group("/kelas")
	{
		router.POST("/add_add", ox.CreateKelas)
		router.POST("/add_siswa", ox.CreateSiswa)
		router.GET("/list_siswa", ox.ListSiswa)
	}
}
