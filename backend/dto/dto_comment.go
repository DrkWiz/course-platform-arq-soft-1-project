package dto

type CommentCreateDto struct {
	Comment string `json:"comment" binding:"required"`
}
