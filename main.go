package main

import (
	"dagi/goRestAPI.com/db"
	"dagi/goRestAPI.com/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()	
    routes.RegisterRoutes(server)

	server.Run(":1234")
}

