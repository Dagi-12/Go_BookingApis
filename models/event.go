// package models

// import "time"

// type Event struct {
// 	ID          int
// 	Name        string `binding:"required"`
// 	Description string `binding:"required"`
// 	Location    string `binding:"required"`
// 	DateTime    time.Time `binding:"required"`
// 	UserId      int
// }
// var events =[]Event{}

// func ( e Event) Save(){
// events = append(events, e)
// }

// func  GetAllEvents()[]Event{
// return events
// }
package models

import (
	"time"

	"dagi/goRestAPI.com/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
	UserId      int       `json:"userId"`
}

var events = []Event{}

func (e Event) Save()error {
	query:= `
	INSERT INTO events(name,description,location,dateTime,user_id)
	VALUES (?,?,?,?,?)
	`
	stmt,err:=db.DB.Prepare(query)
	if err!=nil{
		return  err
	}
	defer stmt.Close()
	result,err:=stmt.Exec(e.Name,e.Description,e.Location,e.DateTime,e.UserId)
	if err!=nil{
     return err
	}
	id,err:=result.LastInsertId()
	
	e.ID=id
	return err
}

func GetAllEvents() []Event {
	return events
}
