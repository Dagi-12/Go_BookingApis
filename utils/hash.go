package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plainPassword string) (string,error) {
	bytes,err:=bcrypt.GenerateFromPassword([]byte(plainPassword),14)
	return string(bytes),err
}