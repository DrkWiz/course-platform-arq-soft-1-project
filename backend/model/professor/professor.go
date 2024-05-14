package model

import usersModel "backend/model/users"

// Domain Classes - "Professor" entity

type Professor struct {
	usersModel.User
	Id_professor int `gorm:"primaryKey"`
}

type Professors []Professor
