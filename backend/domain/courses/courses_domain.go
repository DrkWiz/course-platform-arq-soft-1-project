package courses

type CreateCourseResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ClassAmount int     `json:"classAmount"`
	StartDate   string  `json:"startDate"`
	EndDate     string  `json:"endDate"`
}

type CourseCreateRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ClassAmount int     `json:"classAmount"`
	StartDate   string  `json:"startDate"`
	EndDate     string  `json:"endDate"`
}
