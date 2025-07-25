package middlewares

import (
	"fmt"
	"net/http"

	"dagi/goRestAPI.com/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
token:=context.Request.Header.Get("Authorization")
	if token ==""{
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}
	userId, err:=utils.VerifyToken(token)
	fmt.Println("errorrrrrrrrrrr",err)
if err!=nil{
	context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized"})
	return 
}
context.Set("userId",userId)
context.Next()
}