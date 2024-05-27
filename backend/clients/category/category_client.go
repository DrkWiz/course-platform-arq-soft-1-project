package category

import (
	categoryModel "backend/model/category"

	e "backend/utils/errors"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetCategoryById(id int) categoryModel.Category {

	var category categoryModel.Category

	Db.Where("id_category = ?", id).First(&category)
	log.Debug("Category: ", category)

	return category

}

func CreateCategory(category categoryModel.Category) e.ApiError {
	log.Println("Category to create: ", category)

	err := Db.Save(&category).Error

	if err != nil {
		return e.NewInternalServerApiError("Error creating category", err)
	}
	return nil

}

func CheckCategory(id int) bool {
	var category categoryModel.Category
	Db.Where("id_category = ?", id).First(&category)

	log.Println("Category: ", category.IdCategory)
	return category.IdCategory == 0
}
