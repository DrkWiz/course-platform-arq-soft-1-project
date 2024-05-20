package dto

type CourseMinDto struct {
	IdCourse     int     `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Picture_path string  `json:"picture_path"`
	Start_date   string  `json:"start_date"`
	End_date     string  `json:"end_date"`
}

type CoursesMinDto []CourseMinDto
