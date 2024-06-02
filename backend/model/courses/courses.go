package model

// Domain Classes - "Course" entity
type Course struct {
	IdCourse    int     `gorm:"PRIMARY_KEY;AUTO_INCREMENT;column:id_course"` // Primary key
	Name        string  `gorm:"type:VARCHAR(350);not null"`
	Description string  `gorm:"type:VARCHAR(350);not null"`
	Price       float64 `gorm:"type:float;not null"`
	PicturePath string  `gorm:"type:VARCHAR(350);not null"`
	StartDate   string  `gorm:"type:VARCHAR(350);not null"`
	EndDate     string  `gorm:"type:VARCHAR(350) "`
	IdOwner     int     `gorm:"type:int;not null"`
	IsActive    bool    `gorm:"type:bool;default:true"`
}

type Courses []Course
