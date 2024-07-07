package dto

type CommentCreateDto struct {
	IdUser    int    `json:"id_user" binding:"required"`
	IdCourse  int    `json:"id_course" binding:"required"`
	Comment   string `json:"comment" binding:"required"`
	IsDeleted bool   `json:"is_deleted"`
}
