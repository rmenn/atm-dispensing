package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

const URL string = "root:@tcp(127.0.0.1:3306)/atmdis?parseTime=true"

func main() {
	// Set Gin to production mode
	//gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	router = gin.Default()
	router.Use(DbWare())
	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Initialize the routes
	initializeRoutes()

	// Start serving the application
	router.Run()
}
func Render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}
func DbWare() gin.HandlerFunc {
	db, err := NewDB(URL)
	if err != nil {
		log.Panic(err)
	}
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
