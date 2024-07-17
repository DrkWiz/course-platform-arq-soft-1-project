package dto

// CourseUpdateDto struct

type CourseUpdateDto struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Price       float64          `json:"price"`
	PicturePath string           `json:"picture_path"`
	StartDate   string           `json:"start_date"`
	EndDate     string           `json:"end_date"`
	Categories  CategoriesMinDto `json:"categories"`
	IdOwner     int              `json:"id_owner"`
	IsActive    bool             `json:"is_active"`
}
