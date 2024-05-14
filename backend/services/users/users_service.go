package users

import (
	usersDomain "backend/domain/users"
	"backend/dto"
	usersModel "backend/model/users"
	e "backend/utils/errors"
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

func (s *UsersService) GetStudentById(id int) (dto.StudentMinDto, e.ApiError) {

	var student usersModel.User = GetStudentById(id)
	var studentMinDto dto.StudentMinDto

	if student.Id_user == 0 {
		return studentMinDto, e.NewNotFoundApiError("student not found")
	}

	studentMinDto.IdStudent = student.Id_user
	studentMinDto.Username = student.Username
	studentMinDto.Email = student.Email

	return studentMinDto, nil
}

func Login(request usersDomain.LoginRequest) usersDomain.LoginResponse {

	//validate with db

	return usersDomain.LoginResponse{Token: "nicotroll123"}
}

// GetStudentById method is not working yet because it is not implemented in the database yet

func GetStudentById(id int) usersModel.User {
	return usersModel.User{Id_user: id}
}
