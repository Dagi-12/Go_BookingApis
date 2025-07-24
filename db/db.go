package db

import (
	"database/sql"

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
user_id INTEGER,
FOREIGN KEY(user_id) REFERENCES users(id)
);`

createUser:=`
CREATE TABLE IF NOT EXISTS users(
id INTEGER PRIMARY KEY AUTOINCREMENT,
email TEXT NOT NULL UNIQUE,
password TEXT NOT NULL
)`

createRegistrationTable:=`
CREATE TABLE IF NOT EXISTS registrations(
id INTEGER PRIMARY KEY AUTOINCREMENT,
eventId INTEGER,
userId INTEGER,
FOREIGN KEY(eventId) REFERENCES events(id),
FOREIGN KEY(userId) REFERENCES users(id)
)
`

_,err:=DB.Exec(createUser)
if err!=nil{
	panic("can not create user table")
}

_, err = DB.Exec(createEventsTable)
if err != nil {
	panic("can't create event table")
}

_, err=DB.Exec(createRegistrationTable)
if err != nil {
	panic("can't create registrations table")
}

}
