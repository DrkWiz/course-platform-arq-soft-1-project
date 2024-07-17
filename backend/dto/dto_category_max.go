package dto

type CategoryMaxDto struct {
	IdCategory int    `json:"id_category"`
	Name       string `json:"name"`
}

type CategoriesMaxDto []CategoryMaxDto
