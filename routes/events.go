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
	token:=context.Request.Header.Get("Authorization")
	if token ==""{
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}


	var event models.Event

	err := context.ShouldBindJSON(&event)
	if err != nil {
		fmt.Println("err", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	event.ID = 1
	event.UserId = 1
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
	
 _,err = models.GetEventById(eventId)
    if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't get event"})
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
    if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't get event"})
		return
	}
err= event.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Can not delete event"})
		return
	}
	context.JSON(http.StatusOK,gin.H{"message:":"eventDeleted successfully"})

}
