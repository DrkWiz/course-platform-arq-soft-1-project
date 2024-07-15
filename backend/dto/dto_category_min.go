package dto

type CategoryMinDto struct {
	IdCategory int    `json:"value"`
	Name       string `json:"label"`
}

type CategoriesMinDto []CategoryMinDto
