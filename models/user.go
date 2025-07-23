package models

import (
	"dagi/goRestAPI.com/db"
	"dagi/goRestAPI.com/utils"
	"errors"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query:=`INSERT INTO users(email,password) VALUES (?,?)`
	stmt,err:=db.DB.Prepare(query)
	if err !=nil{
		return err
	}
	hashPass,err:=utils.HashPassword(u.Password)
	
	if err!=nil{
		return err
	}
	defer stmt.Close()
	result,err:=stmt.Exec(u.Email,hashPass)
	if err !=nil{
		return err
	}
	userId,err:=result.LastInsertId()
	u.Id=userId
	return  err


}
func (u *User) ValidateCredentials()error{
	query:=`SELECT id, password FROM users WHERE email = ?`
	row:=db.DB.QueryRow(query,u.Email)
	var retrievedPassword string
	err:=row.Scan(&u.Id ,&retrievedPassword)
	if err!=nil{
		return err
	}
	isPasswordValid:=utils.CheckPasswordHash(u.Password,retrievedPassword)
	if !isPasswordValid{
		return errors.New("credentials is not valid")
	}
  return nil
}