package routes

import (
	"dagi/goRestAPI.com/models"
	"dagi/goRestAPI.com/utils"
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
		return
	}
	err=user.ValidateCredentials()
	if err!=nil{
		fmt.Println(err)
		context.JSON(http.StatusUnauthorized,gin.H{"message":err.Error()})
		return
	}
	token ,err:=utils.GenerateToken(user.Email,user.Id)
	if err!=nil{
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError,gin.H{"message":err.Error()})
		return
	}
	context.JSON(http.StatusOK,gin.H{"message":"Login successful","token":token})
}