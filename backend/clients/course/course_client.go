package course

import (
	courseModel "backend/model/courses"

	e "backend/utils/errors"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetCourseById(id int) courseModel.Course {
	var course courseModel.Course

	Db.Where("id_course = ?", id).First(&course)
	log.Debug("Course: ", course)

	return course
}

func CreateCourse(course courseModel.Course) error {
	err := Db.Create(&course)
	if err != nil {
		return err.Error
	}
	return nil

}

func UpdateCourse(course courseModel.Course) e.ApiError {
	err := Db.Save(&course).Error

	if err != nil {
		return e.NewInternalServerApiError("Error updating course", err)
	}
	return nil
}

func DeleteCourse(id int) error {
	// Ensure the correct usage of the Update method
	err := Db.Model(&courseModel.Course{}).Where("id_course = ?", id).Update("is_active", "0").Error
	if err != nil {
		return err
	}
	return nil
}
