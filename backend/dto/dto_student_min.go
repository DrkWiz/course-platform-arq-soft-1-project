package dto

type UserMinDto struct {
	IdUser   int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	IsAdmin  bool   `json:"is_admin"`
}

type UsersMinDto []UserMinDto
