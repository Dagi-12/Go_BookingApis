// package db

// import (
// 	"database/sql"
// 	"fmt"

// 	_ "github.com/mattn/go-sqlite3"
// )
// var DB *sql.DB
// func InitDB(){
// 	var err error
// 	DB,err = sql.Open("sqlite3","goApi.db")
// 	if err!=nil {
// 		panic("Can not connect to DB")
// 	}

// 	DB.SetMaxOpenConns(10)
// 	DB.SetMaxIdleConns(5)
// 	createTables()
// }

// func createTables(){
// 	createEventsTable := `
// 	CREATE table IF NOT EXISTS events(
// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
// 	name TEXT NOT NULL,
// 	description TEXT NOT NULL,
// 	location TEXT NOT NULL,
// 	dateTime DATETIME NOT NULL,
// 	user_id INTEGER 
// 	)
// 	`
//    _,err:=DB.Exec(createEventsTable)
//    if err!=nil{
// 	fmt.Println("err",err)
// 	panic("cant create event table")
//    }
// }
package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "goApi.db")
	if err != nil {
		panic("Cannot connect to DB")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	);`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		fmt.Println("err", err)
		panic("can't create event table")
	}
}
