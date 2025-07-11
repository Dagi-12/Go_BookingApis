package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)
var DB *sql.DB
func InitDB(){
	var err error
	DB,err = sql.Open("sqlite3","goApi.db")
	if err!=nil {
		panic("Can not connect to DB")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}
