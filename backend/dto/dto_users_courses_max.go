package dto

type UserCourseMaxDto struct {
	IdUser      int              `json:"id_user"`
	IdCourse    int              `json:"id_course"`
	Price       float64          `json:"price"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	PicturePath string           `json:"picture_path"`
	IsActive    bool             `json:"is_active"`
	Categories  CategoriesMaxDto `json:"categories"`
}

type UsersCoursesMaxDto []UserCourseMaxDto
