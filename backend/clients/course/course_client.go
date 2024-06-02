package course

import (
	categoryModel "backend/model/category"
	courseModel "backend/model/courses"

	e "backend/utils/errors"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetCourseById(id int) (courseModel.Course, e.ApiError) {
	var course courseModel.Course

	err := Db.Where("id_course = ?", id).First(&course).Error

	if err != nil {
		return course, e.NewNotFoundApiError("Course not found")
	}

	log.Debug("Course: ", course)

	return course, nil
}

func CreateCourse(course courseModel.Course) e.ApiError {
	err := Db.Create(&course).Error
	if err != nil {
		return e.NewInternalServerApiError("Error creating course", err)
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

func DeleteCourse(id int) e.ApiError {
	// Ensure the correct usage of the Update method
	err := Db.Model(&courseModel.Course{}).Where("id_course = ?", id).Update("is_active", "0").Error
	if err != nil {
		return e.NewInternalServerApiError("Error deleting course", err)
	}
	return nil
}

func CreateCourseCategory(courseId int, categoryId int) e.ApiError {
	courseCategory := courseModel.CourseCategory{IdCourse: courseId, IdCategory: categoryId}
	err := Db.Create(&courseCategory).Error
	if err != nil {
		return e.NewInternalServerApiError("Error creating course category", err)
	}
	return nil
}

func GetCategoriesByCourseId(id int) (categoryModel.Categories, e.ApiError) {
	var categories []categoryModel.Category
	err := Db.Raw("SELECT * FROM categories WHERE id_category IN (SELECT id_category FROM course_categories WHERE id_course = ?)", id).Scan(&categories).Error

	if err != nil {
		return nil, e.NewNotFoundApiError("Categories not found")
	}
	return categories, nil
}

func GetCourses() (courseModel.Courses, e.ApiError) {
	var courses []courseModel.Course
	err := Db.Find(&courses).Error

	if err != nil {
		return nil, e.NewNotFoundApiError("Courses not found")
	}

	return courses, nil
}

func GetOwner(courseId int) (int, e.ApiError) {
	var course courseModel.Course
	err := Db.Where("id_course = ?", courseId).First(&course).Error
	if err != nil {
		return 0, e.NewNotFoundApiError("Course not found")
	}
	return course.IdOwner, nil
}
