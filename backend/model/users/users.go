package model

// Domain Classes - "User" entity
//Parent class of student and professor

type User struct {
	Id_user  int    `gorm:"primaryKey" autoIncrement:"true"`
	Name     string `gorm:"type:varchar(35);not null"`
	Username string `gorm:"type:varchar(35);not null"`
	Password string `gorm:"type:varchar(350);not null"`
	Email    string `gorm:"type:varchar(100);not null"`
	Is_admin bool   `gorm:"type:boolean;not null default:false"`
}

type Users []User
