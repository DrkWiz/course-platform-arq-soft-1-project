package model

// Domain Classes - "Course" entity
type Course struct {
	Id_course    int     `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	Name         string  `gorm:"type:varchar(350);not null"`
	Description  string  `gorm:"type:varchar(350);not null"`
	Price        float64 `gorm:"type:float;not null"`
	Picture_path string  `gorm:"type:varchar(350);not null"`
	Start_date   string  `gorm:"type:varchar(350);not null"`
	End_date     string  `gorm:"type:varchar(350);"`
	Id_user      int     `gorm:"foreignKey:Id_user"`
}

type Courses []Course
