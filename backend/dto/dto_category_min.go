package dto

type CategoryMinDto struct {
	IdCategory int    `json:"id"`
	Name       string `json:"name"`
}

type CategoriesMinDto []CategoryMinDto
