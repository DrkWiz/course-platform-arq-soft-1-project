package model

type CourseCategory struct {
	Id_course   int `gorm:"foreignKey:Id_course"`
	Id_category int `gorm:"foreignKey:Id_category"`
}

type CourseCategoryList []CourseCategory
