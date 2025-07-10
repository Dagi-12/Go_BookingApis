package main

import (
	"net/http"

	"dagi/goRestAPI.com/models"

	"github.com/gin-gonic/gin"
)
func main() {
server:=gin.Default()

server.GET("/events",getEvents)
server.POST("/events",createEvent)

server.Run(":1234")

}
 func getEvents(context *gin.Context)  {
	events:=models.GetAllEvents()
	context.JSON(http.StatusOK,events)
 }
 func createEvent(context *gin.Context){
  var event models.Event
  context.ShouldBindJSON(&event)
  
 }