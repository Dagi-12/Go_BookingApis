package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)
const secretKey="mySafeKey :)(:"
func GenerateToken(email string, userId int64)(string,error) {
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"email":email,
		"userId":userId,
		"exp":time.Now().Add(time.Hour*2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string)(int64, error){
	parsedToken,err:=jwt.Parse(token,func(token *jwt.Token)(interface{},error){
		_,ok:=token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil,errors.New("unexpected signing method")
		}
		return []byte(secretKey),nil
	})
	fmt.Println("in jwt",err)
	if err!= nil {
		return 0, errors.New("could not parse token")
	}
    isValid:=parsedToken.Valid
	if !isValid{
		return 0,errors.New("invalid token")
	}
	claims,ok:=parsedToken.Claims.(jwt.MapClaims)
	if !ok{
		return 0, errors.New("invalid claims")
	}
	// email:=claims["email"].(string)
	userId:=int64(claims["userId"].(float64))
	return userId,nil

}