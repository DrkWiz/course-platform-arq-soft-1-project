package professor

import (
	professorModel "backend/model/professor"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetProfessorById(id int) professorModel.Professor {
	var professor professorModel.Professor

	Db.Where("id = ?", id).First(&professor)
	log.Debug("Professor: ", professor)

	return professor
}
