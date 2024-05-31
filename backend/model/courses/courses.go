package model

// Domain Classes - "Course" entity
type Course struct {
	IdCourse    int     `gorm:"PRIMARY_KEY;AUTO_INCREMENT;column:id_course"` // Primary key
	Name        string  `gorm:"type:VARCHAR(350);not null"`
	Description string  `gorm:"type:VARCHAR(350);not null"`
	Price       float64 `gorm:"type:float;not null"`
	PicturePath string  `gorm:"type:VARCHAR(350);not null"`
	Start_date  string  `gorm:"type:VARCHAR(350);not null"`
	End_date    string  `gorm:"type:VARCHAR(350) "`
	Id_user     int     `gorm:"type:int;not null"`
	IsActive    bool    `gorm:"type:bool;default:true"`
}

type Courses []Course
