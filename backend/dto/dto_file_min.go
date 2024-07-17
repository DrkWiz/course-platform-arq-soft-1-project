package dto

type FileMinDto struct {
	IdFile int `json:"id_file"`
	Name   string `json:"name"`
	Path   string `json:"path"`
}