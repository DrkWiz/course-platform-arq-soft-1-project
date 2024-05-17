package users

import (
	usersClient "backend/clients/users"

	"backend/dto"
	usersModel "backend/model/users"
	e "backend/utils/errors"

	log "github.com/sirupsen/logrus"
)

type usersService struct{}

type usersServiceInterface interface {
	GetStudentById(id int) (dto.StudentMinDto, e.ApiError)
}

var (
	UsersService usersServiceInterface
)

func init() {
	UsersService = &usersService{}
}

func (s *usersService) GetStudentById(id int) (dto.StudentMinDto, e.ApiError) {

	log.Print("GetStudentById: ", id)

	var student usersModel.User = usersClient.GetUserById(id)
	var studentMinDto dto.StudentMinDto

	studentMinDto.IdStudent = student.IdUser
	studentMinDto.Username = student.Username
	studentMinDto.Email = student.Email

	return studentMinDto, nil
}

// GetStudentById method is not working yet because it is not implemented in the database yet

func GetStudentById(id int) usersModel.User {
	return usersModel.User{IdUser: id}
}
