package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func NewDB(dbstring string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dbstring)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

type Atm struct {
	Id        int64     `json:"id"`
	Location  string    `json:"location"`
	Bank      string    `json:"bank"`
	City      string    `json:"city"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func AllAtmsDb(ctx *gin.Context) ([]*Atm, error) {
	db := ctx.MustGet("db").(*sql.DB)
	rows, err := db.Query("SELECT * FROM atms ORDER BY updated_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var atms []*Atm
	for rows.Next() {
		atm := new(Atm)
		err = rows.Scan(&atm.Id,
			&atm.Location,
			&atm.City,
			&atm.Bank,
			&atm.CreatedAt,
			&atm.UpdatedAt)
		if err != nil {
			return nil, err
		}
		atms = append(atms, atm)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return atms, nil
}

func AddAtmDb(ctx *gin.Context) (*Atm, error) {
	db := ctx.MustGet("db").(*sql.DB)
	atm := new(Atm)
	atm.Location = ctx.PostForm("location")
	atm.Bank = ctx.PostForm("bank")
	atm.City = ctx.PostForm("city")
	fmt.Println(atm)
	_, err := db.Exec(`
	INSERT INTO atms(location,bank,city) VALUES(?,?,?)`, &atm.Location, &atm.Bank, &atm.City)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return atm, nil
}

func GetAtmById(ctx *gin.Context) (*Atm, error) {
	atm := new(Atm)
	if ID, err := strconv.Atoi(ctx.Param("id")); err == nil {
		fmt.Println(ID)
		db := ctx.MustGet("db").(*sql.DB)
		row := db.QueryRow("SELECT id, location, city, bank, updated_at, created_at FROM atms WHERE id=?", ID)
		err := row.Scan(&atm.Id, &atm.Location, &atm.Bank, &atm.City, &atm.UpdatedAt, &atm.CreatedAt)
		if err != nil {
			ctx.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		ctx.AbortWithStatus(http.StatusNotFound)
	}
	return atm, nil
}
func ValidateATMDb(ctx *gin.Context) error {
	if Id, err := strconv.Atoi(ctx.Param("id")); err == nil {
		db := ctx.MustGet("db").(*sql.DB)
		_, err := db.Exec(`UPDATE atms SET updated_at=NOW() WHERE id=?`, Id)
		if err != nil {
			return err
		}
	}
	return nil
}
