package model

import usersModel "backend/model/users"

// Domain Classes - "Student" entity
type Student struct {
	usersModel.User
	Id_student int `gorm:"primaryKey"`
}

type Students []Student
