package courses

type Course struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CourseRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
