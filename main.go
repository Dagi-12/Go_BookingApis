// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"dagi/goRestAPI.com/db"
// 	"dagi/goRestAPI.com/models"
// 	"github.com/gin-gonic/gin"
// )
// func main() {
// db.InitDB()
// server:=gin.Default()

// server.GET("/events",getEvents)
// server.POST("/events",createEvent)

// server.Run(":1234")

// }
//
//	 func getEvents(context *gin.Context)  {
//		events:=models.GetAllEvents()
//		context.JSON(http.StatusOK,events)
//	 }
//	 func createEvent(context *gin.Context){
//	  var event models.Event
//	 err:=  context.ShouldBindJSON(&event)
//	 fmt.Println("err",err)
//	  if err!= nil{
//		context.JSON(http.StatusBadRequest,gin.H{"message":"could not parse request data"})
//		return
//	  }
//	  	event.ID=1
//		event.UserId=1
//		event.Save()
//		context.JSON(http.StatusCreated,gin.H{"message":"event created!","event":event})
//	 }
package main

import (
	"fmt"
	"net/http"

	"dagi/goRestAPI.com/db"
	"dagi/goRestAPI.com/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":1234")
}

func getEvents(context *gin.Context) {
	events,err:= models.GetAllEvents()
	if err !=nil {
		context.JSON(http.StatusInternalServerError,gin.H{
			"message":"Could not fetch events,Try Again later",
		})
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)
	if err != nil {
		fmt.Println("err", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	event.ID = 1
	event.UserId = 1
	err =event.Save()
	if err !=nil{
		context.JSON(http.StatusBadRequest,gin.H{"message:":"Can not save file"})
		return 
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event created!", "event": event})
}
