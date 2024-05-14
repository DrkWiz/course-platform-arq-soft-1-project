package course

import (
	courseModel "backend/model/courses"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetCourseById(id int) courseModel.Course {
	var course courseModel.Course

	Db.Where("id = ?", id).First(&course)
	log.Debug("Course: ", course)

	return course
}
