package models

import (
	"university-dean-and-student/app/utility"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID       uint   `json:"id" gorm:"primarykey"`
	Name     string `json:"name" binding:"required"`
	UserName string `json:"user_name" binding:"required"`
	Token    string `json:"token"`
	Password string `json:"password" binding:"required"`
}

func FindUser(username, password string) *User {
	var user User
	result := map[string]string{"user_name": username, "password": password}

	utility.Database().Find(&user, result)
	return &user

}

func (user User) ValidateUser() error {
	validate := validator.New()

	err := validate.Struct(&user)
	if err != nil {
		return err.(validator.ValidationErrors)
	}
	return nil
}

func AuthUser(token string) *User {
	var user User
	result := map[string]string{"token": token}

	utility.Database().Find(&user, result)
	return &user

}
