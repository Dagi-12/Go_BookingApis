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
	UserId      int64     `json:"userId"`
}

// var events = []Event{}

func (e *Event) Save()error {
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

func GetAllEvents() ([]Event,error) {
	query:= "SELECT * FROM events"
	rows,err:=db.DB.Query(query)
	if err!=nil{
		return nil,err
	}
	defer rows.Close()
	var events []Event
	for rows.Next(){
		var event Event
		err:= rows.Scan(&event.ID,&event.Name,&event.Description,&event.Location,&event.DateTime,&event.UserId)
	if err!=nil{
		return nil,err
	}
	events = append(events, event)
	}
	return events,nil
}
func GetEventById(eventId int64)(*Event,error){
	query:="SELECT * FROM events WHERE id = ?"
	row:=db.DB.QueryRow(query,eventId)	
	var event Event
	err:=row.Scan(&event.ID,&event.Name,&event.Description,&event.Location,&event.DateTime,&event.UserId)
	if err!=nil{
		return nil,err
	}
	return &event,nil
}
func (event Event)UpdatedEvent()error{
query := `
UPDATE events
SET name = ?, description = ?, location = ?, dateTime = ?
WHERE id = ?`
	// _, err := db.DB.Exec(query, event.Name, event.Description, event.Location, event.DateTime, event.UserId, event.ID)
    stmt,err:=db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_,err =stmt.Exec(event.Name,event.Description,event.Location,event.DateTime,event.ID)
	return  err
}
func (event *Event)DeleteEvent()error{
	query :=`
	DELETE FROM events
	WHERE id = ?
	`
	stmt,err:=db.DB.Prepare(query)	
		if err != nil {
		return err
	}
	defer stmt.Close()
	_,err =stmt.Exec(event.ID)
	return err

}
func(e Event) Register(userId int64)error{
	query:="INSERT INTO registrations(eventId,userId)VALUES(?,?)"
	stmt,err:=db.DB.Prepare(query)
	if err!=nil{
		return err
	}
	defer stmt.Close()
	_,err=stmt.Exec(e.ID,userId)
	return err
}
func (e Event) CancelRegistration(userId int64)error{
query:=`
DELETE FROM registrations WHERE eventId = ? AND userId =?
`
stmt,err:=db.DB.Prepare(query)
	if err!=nil{
		return err
	}
	defer stmt.Close()
	_,err=stmt.Exec(e.ID,userId)
	return err
}