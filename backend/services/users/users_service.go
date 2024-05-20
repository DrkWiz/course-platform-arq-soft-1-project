package users

import (
	usersClient "backend/clients/users"

	"backend/dto"
	usersModel "backend/model/users"
	e "backend/utils/errors"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
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

// Password hashing
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//Create user

func CreateUser(user dto.UserCreateDto) error {
	//TODO: IMPLEMENT PASSWORD HASHING
	hashPassword, err := HashPassword(user.Password)

	if err != nil {
		return err
	}

	userToCreate := usersModel.User{Name: user.Name, Username: user.Username, Email: user.Email, Password: hashPassword}

	err = usersClient.CreateUser(userToCreate)
	if err != nil {
		return err
	}
	return nil
}
