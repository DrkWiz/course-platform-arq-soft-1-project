package dto

type CategoryTokenDto struct {
	IdCategory int    `json:"id_category"`
	Name       string `json:"name"`
}

type CategoriesTokenDto []CategoryTokenDto

type UserCourseTokenDto struct {
	IdUser      int                `json:"id_user"`
	IdCourse    int                `json:"id_course"`
	Price       float64            `json:"price"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	PicturePath string             `json:"picture_path"`
	IsActive    bool               `json:"is_active"`
	Categories  CategoriesTokenDto `json:"categories"`
}

type UsersCoursesTokenDto []UserCourseTokenDto
