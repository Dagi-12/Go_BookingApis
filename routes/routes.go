package routes

import (
	"dagi/goRestAPI.com/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// server.POST("/events",middlewares.Authenticate, createEvent)
	// server.PUT("/events/:eventId",middlewares.Authenticate,updateEvent)
	// server.DELETE("/events/:eventId",middlewares.Authenticate,deleteEvent)
	server.GET("/events", getEvents)
	server.GET("/events/:eventId", getSingleEvent)
	server.POST("/users/signUp",signUp)
	server.POST("/users/signIn",signIn)

	authenticated:=server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:eventId",updateEvent)
	authenticated.DELETE("/events/:eventId",deleteEvent)
	authenticated.POST("/events/:eventId/register",registerForEvent)
	authenticated.DELETE("/events/:eventId/cancel",cancelRegistration)
}