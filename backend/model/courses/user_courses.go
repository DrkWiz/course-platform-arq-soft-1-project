package model

type UserCourses struct {
	Id_user   int     `gorm:"foreignKey:Id_user"`
	Id_course int     `gorm:"foreignKey:Id_course"`
	Rating    float64 `gorm:"type:float;"`
	Comment   string  `gorm:"type:varchar(350);"`
}

type UserCoursesList []UserCourses
