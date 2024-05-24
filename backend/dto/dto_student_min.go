package dto

type UserMinDto struct {
	IdUser   int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UsersMinDto []UserMinDto
