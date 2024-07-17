package dto

type CourseMinDto struct {
	IdCourse    int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	PicturePath string  `json:"picture_path"`
	StartDate   string  `json:"start_date"`
	EndDate     string  `json:"end_date"`
	IsActive    bool    `json:"is_active"`
}

type CoursesMinDto []CourseMinDto
