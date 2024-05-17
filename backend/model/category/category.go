package model

type Category struct {
	IdCategory int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	Name       string `gorm:"type:varchar(35);not null"`
}

type Categories []Category
