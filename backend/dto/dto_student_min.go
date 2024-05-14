package dto

type StudentMinDto struct {
	IdStudent int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

type StudentsMinDto []StudentMinDto
