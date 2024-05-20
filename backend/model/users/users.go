package model

// Domain Classes - "User" entity
//Parent class of student and professor

type User struct {
	IdUser   int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	Name     string `gorm:"type:varchar(35);not null"`
	Username string `gorm:"type:varchar(35);not null"`
	Password string `gorm:"type:varchar(350);not null"`
	Email    string `gorm:"type:varchar(100);not null"`
	IsAdmin  bool   `gorm:"type:boolean;not null default:false"`
}

type Users []User
