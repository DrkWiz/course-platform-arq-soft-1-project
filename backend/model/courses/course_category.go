package model

type CourseCategory struct {
	IdCourse   int `gorm:"type:int;not null"`
	IdCategory int `gorm:"type:int;not null"`
}

type CourseCategoryList []CourseCategory
