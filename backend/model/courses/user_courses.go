package model

type UserCourses struct {
	Id_user   int     `gorm:"type:int;not null"`
	Id_course int     `gorm:"type:int;not null"`
	Rating    float64 `gorm:"type:float;"`
	Comment   string  `gorm:"type:varchar(350);"`
}

type UserCoursesList []UserCourses
