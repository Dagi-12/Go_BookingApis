package routes

import (
	"net/http"
	"strconv"

	"dagi/goRestAPI.com/models"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
 userId:=context.GetInt64("userId")
 eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "couldn't parse event id"})
		return
	}
event,err:=models.GetEventById(eventId)

if err != nil {
	context.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't get event"})
	return
 }
err=event.Register(userId)
if err!=nil{
context.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't register for event"})
	return
}
context.JSON(http.StatusCreated, gin.H{"message": "successfully registered for event"})

}
func cancelRegistration(context *gin.Context) {

}