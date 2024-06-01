package model

type UserCourses struct {
	IdUser   int     `gorm:"type:int;not null"`
	IdCourse int     `gorm:"type:int;not null"`
	Rating   float64 `gorm:"type:float;"`
	Comment  string  `gorm:"type:varchar(350);"`
}

type UserCoursesList []UserCourses
