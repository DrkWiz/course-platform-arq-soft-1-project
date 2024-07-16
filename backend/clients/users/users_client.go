package users

import (
	userModel "backend/model/users"

	e "backend/utils/errors"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

// Devuelve un usuario con id
func GetUserById(id int) (userModel.User, e.ApiError) {
	var user userModel.User

	err := Db.Where("id_user = ?", id).First(&user).Error
	if err != nil {
		return user, e.NewNotFoundApiError("User not found")
	}
	log.Print("User: ", user)
	return user, nil
}

// Crea un usuario
func CreateUser(user userModel.User) e.ApiError {
	err := Db.Create(&user).Error
	if err != nil {
		return e.NewInternalServerApiError("Error creating user", err)
	}
	return nil
}

// Revisa si un username ya existe.
func CheckUsername(username string) bool {
	var user userModel.User
	Db.Where("username = ?", username).First(&user)
	return user.Username != ""
}

// Revisa si un email ya existe.
func CheckEmail(email string) bool {
	var user userModel.User
	Db.Where("email = ?", email).First(&user)
	return user.Username != ""
}

// Revisa una inscripcion.
func CheckUserCourse(idUser int, idCourse int) bool {
	var userCourse userModel.UserCourses
	Db.Where("id_user = ? AND id_course = ?", idUser, idCourse).First(&userCourse)
	return userCourse.IdUser != 0
}

// Devuelve un usuario con username
func GetUserByUsername(username string) (userModel.User, e.ApiError) {
	var user userModel.User
	err := Db.Where("username = ?", username).First(&user).Error

	if err != nil {
		return user, e.NewNotFoundApiError("User not found")
	}

	return user, nil
}

// Devuelve las inscripciones de un usuario
func GetUserCourses(id int) ([]userModel.UserCourses, e.ApiError) {
	var usercourses []userModel.UserCourses
	err := Db.Where("id_user = ?", id).Find(&usercourses).Error

	if err != nil {
		return nil, e.NewInternalServerApiError("Error getting user-courses", err)
	}

	return usercourses, nil
}

// Agrega una inscripcion a la db.
func AddUserCourse(userCourse userModel.UserCourses) e.ApiError {
	err := Db.Create(&userCourse).Error
	if err != nil {
		return e.NewInternalServerApiError("Error creating user-course", err)
	}
	return nil
}

// Elimina una inscripcion de la db.
func RemoveUserCourse(idUser int, idCourse int) e.ApiError {
	var userCourse userModel.UserCourses

	Db.Where("id_user = ? AND id_course = ?", idUser, idCourse).First(&userCourse)
	if userCourse.IdUser == 0 {
		return e.NewBadRequestApiError("User is not subscribed to course")
	}
	Db.Exec("DELETE FROM user_courses WHERE id_user = ? AND id_course = ?", idUser, idCourse)
	return nil
}
