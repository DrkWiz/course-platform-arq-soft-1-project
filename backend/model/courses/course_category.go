package model

type CourseCategory struct {
	IdCourse   int `gorm:"foreignKey:Id_course"`
	IdCategory int `gorm:"foreignKey:Id_category"`
}

type CourseCategoryList []CourseCategory
