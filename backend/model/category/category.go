package model

type Category struct {
	IdCategory int    `gorm:"PRIMARY_KEY"`
	Name       string `gorm:"type:varchar(35);not null"`
}

type Categories []Category
