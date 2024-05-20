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
	GetUserById(id int) (dto.UserMinDto, e.ApiError)
}

var (
	UsersService usersServiceInterface
)

func init() {
	UsersService = &usersService{}
}

func (s *usersService) GetUserById(id int) (dto.UserMinDto, e.ApiError) {

	log.Print("GetUserById: ", id)

	var user usersModel.User = usersClient.GetUserById(id)
	var UserMinDto dto.UserMinDto

	UserMinDto.IdUser = user.IdUser
	UserMinDto.Username = user.Username
	UserMinDto.Email = user.Email

	return UserMinDto, nil
}

//Create user

func CreateUser(user dto.UserCreateDto) error {
	//TODO: IMPLEMENT PASSWORD HASHING
	hashPassword := user.Password
	userToCreate := usersModel.User{Name: user.Name, Username: user.Username, Email: user.Email, Password: hashPassword}

	err := usersClient.CreateUser(userToCreate)
	if err != nil {
		return err
	}
	return nil
}
