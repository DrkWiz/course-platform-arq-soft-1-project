package users

import (
	userModel "backend/model/users"
	userCoursesModel "backend/model/users/user_courses"

	e "backend/utils/errors"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetUserById(id int) userModel.User {
	var user userModel.User

	Db.Where("id_user = ?", id).First(&user)
	log.Print("User: ", user)

	return user
}

func CreateUser(user userModel.User) e.ApiError {
	err := Db.Create(&user).Error
	if err != nil {
		return e.NewInternalServerApiError("Error creating user", err)
	}
	return nil
}

func CheckUsername(username string) bool {
	var user userModel.User
	Db.Where("username = ?", username).First(&user)
	return user.Username != ""
}

func CheckEmail(email string) bool {
	var user userModel.User
	Db.Where("email = ?", email).First(&user)
	return user.Username != ""
}

func GetUserByUsername(username string) userModel.User {
	var user userModel.User
	Db.Where("username = ?", username).First(&user)
	return user
}

//UserCourses

func GetUserCourses(id int) []userCoursesModel.UserCourses {
	var usercourses []userCoursesModel.UserCourses
	Db.Where("id_user = ?", id).Find(&usercourses)
	return usercourses
}
