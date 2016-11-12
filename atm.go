package main

func initializeRoutes() {
	router.POST("/atm/add", AddAtm)
	router.POST("/atm/:uuid/validate", ValidateATM)
	router.GET("/", ShowAll)
	router.GET("/:query", SearchCity)
}
