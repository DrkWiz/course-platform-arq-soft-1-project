package dto

type UserLoginDto struct {
	Username string `json:"username"`
	Password string `json:"password" binding:"required"`
}
