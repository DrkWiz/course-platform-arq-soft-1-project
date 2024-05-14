package model

// Domain Classes - "User" entity
//Parent class of student and professor

type User struct {
	Id_user  int    `gorm:"primaryKey"`
	Username string `gorm:"type:varchar(350);not null"`
	Password string `gorm:"type:varchar(350);not null"`
	Email    string `gorm:"type:varchar(350);not null"`
	Is_admin bool   `gorm:"type:boolean;not null"`
}
