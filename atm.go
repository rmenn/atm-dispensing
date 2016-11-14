package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func ShowAll(ctx *gin.Context) {
	atms, err := AllAtmsDb(ctx)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(404, gin.H{"error": "None found"})
		return
	}
	Render(ctx, gin.H{
		"title":   "ATMS",
		"payload": atms},
		"atms.html")
}

func ShowAtm(ctx *gin.Context) {
	atm, err := GetAtmById(ctx)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(404, gin.H{"error": "Not Found"})
		return
	}
	Render(ctx, gin.H{
		"title":   atm.Location,
		"payload": atm},
		"atm-view.html")
}
func AddPage(ctx *gin.Context) {
	Render(ctx, gin.H{
		"title": "Add ATM"},
		"addatm.html")
}
func AddAtm(ctx *gin.Context) {
	_, err := AddAtmDb(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal error"})
		return
	}
	Render(ctx, gin.H{
		"title": "Add ATM"},
		"success.html")
}
func ValidateATM(ctx *gin.Context) {
	err := ValidateATMDb(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal error"})
		return
	}
	Render(ctx, gin.H{
		"title": "Validate ATM"},
		"success.html")
}
func SearchCity(ctx *gin.Context) {}
