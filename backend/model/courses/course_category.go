package model

type CourseCategory struct {
	Id_course   int `gorm:"type:int;not null"`
	Id_category int `gorm:"type:int;not null"`
}

type CourseCategoryList []CourseCategory
