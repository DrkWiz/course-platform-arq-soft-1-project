package dto

type CourseMaxDto struct {
	IdCourse    int                `json:"id"`
	Owner       int                `json:"owner"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Price       float64            `json:"price"`
	PicturePath string             `json:"picture_path"`
	StartDate   string             `json:"start_date"`
	EndDate     string             `json:"end_date"`
	IsActive    bool               `json:"is_active"`
	Categories  CategoriesTokenDto `json:"categories"`
}

type CoursesMaxDto []CourseMaxDto
