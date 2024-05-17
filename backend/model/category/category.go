package model

type Category struct {
	Id_category int    `gorm:"primaryKey" autoIncrement:"true"`
	Name        string `gorm:"type:varchar(35);not null"`
}

type Categories []Category
