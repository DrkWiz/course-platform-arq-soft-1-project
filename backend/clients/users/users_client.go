package users

import (
	userModel "backend/model/users"

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

func CheckUserCourse(idUser int, idCourse int) bool {
	var userCourse userModel.UserCourses
	Db.Where("id_user = ? AND id_course = ?", idUser, idCourse).First(&userCourse)
	return userCourse.IdUser != 0
}

func GetUserByUsername(username string) (userModel.User, e.ApiError) {
	var user userModel.User
	err := Db.Where("username = ?", username).First(&user).Error

	if err != nil {
		return user, e.NewNotFoundApiError("User not found")
	}

	return user, nil
}

//UserCourses

func GetUserCourses(id int) ([]userModel.UserCourses, e.ApiError) {
	var usercourses []userModel.UserCourses
	err := Db.Where("id_user = ?", id).Find(&usercourses).Error

	if err != nil {
		return nil, e.NewInternalServerApiError("Error getting user-courses", err)
	}

	return usercourses, nil
}

// Agregar un user-course

func AddUserCourse(userCourse userModel.UserCourses) e.ApiError {
	err := Db.Create(&userCourse).Error
	if err != nil {
		return e.NewInternalServerApiError("Error creating user-course", err)
	}
	return nil
}

// Eliminar un user-course

func RemoveUserCourse(idUser int, idCourse int) e.ApiError {
	var userCourse userModel.UserCourses

	Db.Where("id_user = ? AND id_course = ?", idUser, idCourse).First(&userCourse)
	if userCourse.IdUser == 0 {
		return e.NewBadRequestApiError("User is not subscribed to course")
	}
	Db.Exec("DELETE FROM user_courses WHERE id_user = ? AND id_course = ?", idUser, idCourse)
	return nil
}
