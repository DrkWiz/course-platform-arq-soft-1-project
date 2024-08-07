package dto

type UserCourseMinDto struct {
	IdUser   int     `json:"id_user"`
	IdCourse int     `json:"id_course"`
	Rating   float64 `json:"rating"`
	Comment  string  `json:"comment"`
	IsActive bool    `json:"is_active"`
}

type UserCoursesMinDto []UserCourseMinDto
