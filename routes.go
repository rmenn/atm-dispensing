package main

func initializeRoutes() {
	router.GET("/atm/add", AddPage)
	router.POST("/atm/add", AddAtm)
	router.GET("/atm/view/:id", ShowAtm)
	router.POST("/atm/validate/:id", ValidateATM)
	router.GET("/", ShowAll)
	router.GET("/search/:query", SearchCity)
}
