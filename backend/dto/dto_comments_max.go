package dto

type CommentMaxDto struct {
	Username string  `json:"username"`
	Comment  string  `json:"comment"`
	Rating   float64 `json:"rating"`
}

type CommentsMaxDto []CommentMaxDto
