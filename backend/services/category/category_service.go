package category

import (
	categoryClient "backend/clients/category"

	"backend/dto"

	categoryModel "backend/model/category"

	e "backend/utils/errors"

	log "github.com/sirupsen/logrus"
)

type categoryService struct{}

type categoryServiceInterface interface {
	GetCategoryById(id int) (dto.CategoryMinDto, e.ApiError)
}

var (
	CategoryService categoryServiceInterface
)

func init() {
	CategoryService = &categoryService{}
}

func (s *categoryService) GetCategoryById(id int) (dto.CategoryMinDto, e.ApiError) {

	log.Print("GetCategoryById: ", id)

	var category categoryModel.Category = categoryClient.GetCategoryById(id)
	var CategoryMinDto dto.CategoryMinDto

	CategoryMinDto.IdCategory = category.IdCategory
	CategoryMinDto.Name = category.Name

	return CategoryMinDto, nil
}

func CreateCategory(category dto.CategoryCreateDto) error {

	categoryToCreate := categoryModel.Category{Name: category.Name}

	err := categoryClient.CreateCategory(categoryToCreate)

	if err != nil {
		return err
	}
	return nil
}
