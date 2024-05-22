package category

import (
	categoryModel "backend/model/category"

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

func CreateCategory(category categoryModel.Category) error {

	err := Db.Create(&category)

	if err != nil {
		return err.Error
	}
	return nil

}
