package routes

import (
	"dagi/goRestAPI.com/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch events,Try Again later",
		})
	}
	context.JSON(http.StatusOK, events)
}

func getSingleEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "couldn't parse event id"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Cant get the specified event"})
		return
	}
	context.JSON(http.StatusOK, event)

}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		fmt.Println("err", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}
	
	userId:=context.GetInt64("userId")
	event.ID = 1
	event.UserId = userId
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message:": "Can not save file"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event created!", "event": event})
}

func updateEvent(context *gin.Context){
 eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "couldn't parse event id"})
		return
	}
     userId:=context.GetInt64("userId")	
     event,err := models.GetEventById(eventId)


    if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't get event"})
		return
	}
	if event.UserId != userId{
		context.JSON(http.StatusUnauthorized, gin.H{"message": "couldn't update event you  are not authorized"})
		return
	}
 
 var updatedEvent models.Event
 err = context.ShouldBindJSON(&updatedEvent)
 if err!=nil{
	context.JSON(http.StatusBadRequest, gin.H{"message": "couldn't parse request data"})
	return
 }
 updatedEvent.ID = eventId
 err=updatedEvent.UpdatedEvent()
 
 if err!=nil{
	context.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't update event"})
	return
 }
 context.JSON(http.StatusOK,gin.H{"message":"Event updated successfully"})

}

func deleteEvent(context *gin.Context){
 eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "couldn't parse event id"})
		return
	}
 event,err := models.GetEventById(eventId)
userId:=context.GetInt64("userId")	
    if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't get event"})
		return
	}
	if event.UserId != userId{
		context.JSON(http.StatusUnauthorized, gin.H{"message": "couldn't update event you are not authorized"})
		return
	}
err= event.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Can not delete event"})
		return
	}
	context.JSON(http.StatusOK,gin.H{"message:":"eventDeleted successfully"})

}
