package dto

type CourseCreateDto struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	PicturePath string  `json:"picture_path"`
	StartDate   string  `json:"start_date"`
	EndDate     string  `json:"end_date"`
	IdOwner     int     `json:"id_owner"`
}
