package student

import (
	studentModel "backend/model/student"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetStudentById(id int) studentModel.Student {
	var student studentModel.Student

	Db.Where("id = ?", id).First(&student)
	log.Debug("Student: ", student)

	return student
}
