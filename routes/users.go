package routes

import (
	"dagi/goRestAPI.com/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
var user models.User
err := context.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println("err", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}
err=user.Save()
  if err!=nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save user"})

  }
  context.JSON(http.StatusOK,gin.H{"message":"User Created Successfully"})
}

func signIn (context *gin.Context){
	var user models.User
	err:= context.ShouldBindJSON(&user)
	if err!=nil{
		fmt.Println(err)
		context.JSON(http.StatusBadRequest,gin.H{"message:":"could not parse request body"})
	}
}