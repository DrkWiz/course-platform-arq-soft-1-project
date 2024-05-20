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
	GetUserById(id int) (dto.StudentMinDto, e.ApiError)
}

var (
	UsersService usersServiceInterface
)

func init() {
	UsersService = &usersService{}
}

func (s *usersService) GetUserById(id int) (dto.StudentMinDto, e.ApiError) {

	log.Print("GetUserById: ", id)

	var student usersModel.User = usersClient.GetUserById(id)
	var studentMinDto dto.StudentMinDto

	studentMinDto.IdStudent = student.IdUser
	studentMinDto.Username = student.Username
	studentMinDto.Email = student.Email

	return studentMinDto, nil
}

// GetUseById method is not working yet because it is not implemented in the database yet

func GetUserById(id int) usersModel.User {
	return usersModel.User{IdUser: id}
}
