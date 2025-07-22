package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/events", createEvent)
	server.GET("/events", getEvents)
	server.GET("/events/:eventId", getSingleEvent)
	server.PUT("/events/:eventId",updateEvent)
	server.DELETE("/events/:eventId",deleteEvent)
	server.POST("/users/signUp",signUp)
	server.POST("/users/signIn",signIn)
}