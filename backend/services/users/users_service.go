package users

import (
	courseClient "backend/clients/course"
	usersClient "backend/clients/users"
	"time"

	"backend/dto"
	usersModel "backend/model/users"
	e "backend/utils/errors"

	jwt "github.com/golang-jwt/jwt/v4"

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

func CreateUser(user dto.UserCreateDto) e.ApiError {

	// Check if username or email already exists
	if usersClient.CheckUsername(user.Username) {
		err := e.NewBadRequestApiError("Username already exists")
		return err
	}
	if usersClient.CheckEmail(user.Email) {
		err := e.NewBadRequestApiError("Email is already used")
		return err
	}

	hashPassword, err := HashPassword(user.Password)

	if err != nil {
		return e.NewInternalServerApiError("Error hashing password", err)
	}

	userToCreate := usersModel.User{Name: user.Name, Username: user.Username, Email: user.Email, Password: hashPassword}

	err1 := usersClient.CreateUser(userToCreate)
	if err1 != nil {
		return err1
	}
	return nil
}

// Login user

func Login(user dto.UserLoginDto) (string, e.ApiError) {
	userToLogin := usersClient.GetUserByUsername(user.Username)

	if userToLogin.IdUser == 0 {
		return "", e.NewBadRequestApiError("Invalid username")
	}

	err := bcrypt.CompareHashAndPassword([]byte(userToLogin.Password), []byte(user.Password))

	if err != nil {
		return "", e.NewBadRequestApiError("Invalid password")
	}

	token, err1 := createToken(userToLogin.IdUser)

	if err1 != nil {
		return "", err1
	}

	return token, nil
}

// create token
func createToken(id int) (string, e.ApiError) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("fremico"))

	if err != nil {
		return "", e.NewInternalServerApiError("Error creating token", err)
	}

	return tokenString, nil
}

func validateToken(tokenString string) (int, e.ApiError) {

	log.Print("Token: ", tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("fremico"), nil
	})

	if err != nil {
		return 0, e.NewBadRequestApiError("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	log.Print("Claims: ", claims)

	if !ok {
		return 0, e.NewBadRequestApiError("Invalid token")
	}

	if claims["exp"].(float64) < float64(time.Now().Unix()) {
		return 0, e.NewBadRequestApiError("Token expired")
	}

	id := int(claims["id"].(float64))

	return id, nil
}

func GetUsersByToken(token string) (dto.UserMinDto, e.ApiError) {
	id, err := validateToken(token)

	if err != nil {
		return dto.UserMinDto{}, err
	}

	user, err1 := UsersService.GetUserById(id)

	if err1 != nil {
		return dto.UserMinDto{}, err1
	}

	return user, nil
}

// usercourses
func GetUserCourses(id int) (dto.UserCoursesMinDto, e.ApiError) {
	usercourses := usersClient.GetUserCourses(id)

	var userCoursesMinDto dto.UserCoursesMinDto

	for _, usercourse := range usercourses {
		var userCourseMinDto dto.UserCourseMinDto

		// Mapping fields from UserCourses to UserCourseMinDto
		userCourseMinDto.IdUser = usercourse.IdUser
		userCourseMinDto.IdCourse = usercourse.IdCourse
		// Converting float64 to int for Rating if needed
		userCourseMinDto.Rating = int(usercourse.Rating)
		userCourseMinDto.Comment = usercourse.Comment

		userCourseMinDto.IsActive = courseClient.GetCourseById(usercourse.IdCourse).IsActive
		// Appending the mapped dto to the slice
		userCoursesMinDto = append(userCoursesMinDto, userCourseMinDto)
	}

	return userCoursesMinDto, nil
}
