package category

import (
	categoryModel "backend/model/category"

	e "backend/utils/errors"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetCategoryById(id int) (categoryModel.Category, e.ApiError) {

	var category categoryModel.Category

	err := Db.Where("id_category = ?", id).First(&category).Error

	if err != nil {
		return category, e.NewNotFoundApiError("Category not found")
	}
	log.Debug("Category: ", category)

	return category, nil
}

func CreateCategory(category categoryModel.Category) error {

	err := Db.Create(&category)

	if err != nil {
		return err.Error
	}
	return nil

}

func GetCategories() (categoryModel.Categories, e.ApiError) {

	var categories []categoryModel.Category

	err := Db.Find(&categories).Error

	if err != nil {
		return nil, e.NewInternalServerApiError("Error getting categories", err)
	}
	log.Debug("Categories: ", categories)

	return categories, nil
}
