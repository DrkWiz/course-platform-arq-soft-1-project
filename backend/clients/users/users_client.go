package users

import (
	userModel "backend/model/users"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetUserById(id int) userModel.User {
	var user userModel.User

	Db.Where("user_id = ?", id).First(&user)
	log.Print("User: ", user)

	return user
}
