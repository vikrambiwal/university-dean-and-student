package models

import (
	"fmt"
	"time"
	"university-dean-and-student/app/utility"

	"github.com/go-playground/validator/v10"
)

type Session struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	StudentId int64     `json:"student_id" binding:"required"`
	DeanId    int64     `json:"dean_id" binding:"required"`
	Date      time.Time `json:"date" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
	UpdatedAt time.Time `json:"updated_at" binding:"required"`
}

func (session Session) ValidateSession() error {
	validate := validator.New()

	err := validate.Struct(&session)
	if err != nil {
		return err.(validator.ValidationErrors)
	}
	return nil
}

func SessionListForStudents(userId int64) []Session {
	var sessions []Session

	condition := "date >= ?"
	result := map[string]interface{}{"student_id": fmt.Sprint(userId)}
	utility.Database().Where(condition, time.Now()).Find(&sessions, result)
	return sessions
}

func SessionListForDean(userId int64) []Session {
	var sessions []Session

	condition := "date >= ?"
	result := map[string]string{"dean_id": fmt.Sprint(userId)}
	utility.Database().Where(condition, time.Now()).Find(&sessions, result)
	return sessions
}
