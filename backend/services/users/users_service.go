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
	HashPassword(password string) (string, e.ApiError)
	CreateUser(user dto.UserCreateDto) e.ApiError
	Login(user dto.UserLoginDto) (string, e.ApiError)
	createToken(id int) (string, e.ApiError)
	ValidateToken(tokenString string) (int, e.ApiError)
	GetUsersByToken(token string) (dto.UserMinDto, e.ApiError)
	GetUserCourses(id int) (dto.UserCoursesMinDto, e.ApiError)
	AddUserCourse(idCourse int, token string) e.ApiError
	GetUserCoursesByToken(token string) (dto.UsersCoursesMaxDto, e.ApiError)
	CheckAdmin(token string) (bool, e.ApiError)
	UnsubscribeUserCourse(courseId int, token string) e.ApiError
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
	UserMinDto.Name = user.Name
	UserMinDto.Email = user.Email
	UserMinDto.IsAdmin = user.IsAdmin

	return UserMinDto, nil
}

// Password hashing
func (s *usersService) HashPassword(password string) (string, e.ApiError) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return "", e.NewInternalServerApiError("Error hashing password", err)
	}

	return string(bytes), nil
}

//Create user

func (s *usersService) CreateUser(user dto.UserCreateDto) e.ApiError {

	// Check if username or email already exists
	if usersClient.CheckUsername(user.Username) {
		err := e.NewBadRequestApiError("Username already exists")
		return err
	}
	if usersClient.CheckEmail(user.Email) {
		err := e.NewBadRequestApiError("Email is already used")
		return err
	}

	hashPassword, err := s.HashPassword(user.Password)

	if err != nil {
		return err
	}

	userToCreate := usersModel.User{Name: user.Name, Username: user.Username, Email: user.Email, Password: hashPassword, IsAdmin: false}

	err1 := usersClient.CreateUser(userToCreate)
	if err1 != nil {
		return err1
	}
	return nil
}

// Login user

func (s *usersService) Login(user dto.UserLoginDto) (string, e.ApiError) {
	userToLogin, err := usersClient.GetUserByUsername(user.Username)

	if err != nil {
		return "", err
	}

	if userToLogin.IdUser == 0 {
		return "", e.NewBadRequestApiError("Invalid username")
	}

	err1 := bcrypt.CompareHashAndPassword([]byte(userToLogin.Password), []byte(user.Password))

	if err1 != nil {
		return "", e.NewBadRequestApiError("Invalid password")
	}

	token, err := s.createToken(userToLogin.IdUser)

	if err != nil {
		return "", err
	}

	return token, nil
}

// create token
func (s *usersService) createToken(id int) (string, e.ApiError) {
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

func (s *usersService) ValidateToken(tokenString string) (int, e.ApiError) {
	log.Print("Token: ", tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("fremico"), nil
	})

	if err != nil {
		return 0, e.NewBadRequestApiError("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	log.Print("Claims: ", claims)

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return 0, e.NewBadRequestApiError("Token expired")
	}

	if !ok || !token.Valid {
		return 0, e.NewBadRequestApiError("Invalid token")
	}

	id := int(claims["id"].(float64))

	return id, nil
}

func (s *usersService) GetUsersByToken(token string) (dto.UserMinDto, e.ApiError) {
	id, err := s.ValidateToken(token)

	if err != nil {
		return dto.UserMinDto{}, err
	}

	user, err1 := s.GetUserById(id)

	if err1 != nil {
		return dto.UserMinDto{}, err1
	}

	return user, nil
}

// usercourses
func (s *usersService) GetUserCourses(id int) (dto.UserCoursesMinDto, e.ApiError) {
	userCourses, err := usersClient.GetUserCourses(id)

	if err != nil {
		return nil, err
	}

	// Declaramos la variable a devolver
	var userCoursesMinDto dto.UserCoursesMinDto

	for _, usercourse := range userCourses {
		var userCourseMinDto dto.UserCourseMinDto

		// Mapping fields from UserCourses to UserCourseMinDto
		userCourseMinDto.IdUser = usercourse.IdUser
		userCourseMinDto.IdCourse = usercourse.IdCourse
		// Converting float64 to int for Rating if needed
		userCourseMinDto.Rating = usercourse.Rating
		userCourseMinDto.Comment = usercourse.Comment

		// Conseguimos el curso de este usercourse
		course, err := courseClient.GetCourseById(usercourse.IdCourse)

		if err != nil {
			return nil, err
		}

		userCourseMinDto.IsActive = course.IsActive
		// Appending the mapped dto to the slice
		userCoursesMinDto = append(userCoursesMinDto, userCourseMinDto)
	}

	return userCoursesMinDto, nil
}

// Inscribir usuario a curso

func (s *usersService) AddUserCourse(idCourse int, token string) e.ApiError {
	idUser, err := s.ValidateToken(token)

	if err != nil {
		return err
	}

	// Check if user is already enrolled in the course
	if usersClient.CheckUserCourse(idUser, idCourse) {
		err := e.NewBadRequestApiError("User is already enrolled in this course")
		return err
	}

	userCourseToAdd := usersModel.UserCourses{IdUser: idUser, IdCourse: idCourse, Rating: 0, Comment: ""}
	err1 := usersClient.AddUserCourse(userCourseToAdd)

	if err1 != nil {
		return err
	}

	return nil
}

// Falta manejo de errores.

func (s *usersService) GetUserCoursesByToken(token string) (dto.UsersCoursesMaxDto, e.ApiError) {
	id, err := s.ValidateToken(token) // Valida el token y devuelve id.

	if err != nil {
		return nil, err
	}

	userCourses, err := s.GetUserCourses(id) // Devuelve array de tabla user_courses.

	if err != nil {
		return nil, err
	}

	userCoursesDto := dto.UsersCoursesMaxDto{} // Variable a llenar y retornar

	for _, userCourse := range userCourses { // Recorre el array de user_courses.
		var userMaxCourseDto dto.UserCourseMaxDto

		userMaxCourseDto.IdUser = userCourse.IdUser
		userMaxCourseDto.IdCourse = userCourse.IdCourse

		// Obtenemos el curso
		tempCourse, err := courseClient.GetCourseById(userCourse.IdCourse)

		if err != nil {
			return nil, err
		}

		userMaxCourseDto.Price = tempCourse.Price
		userMaxCourseDto.Name = tempCourse.Name
		userMaxCourseDto.Description = tempCourse.Description
		userMaxCourseDto.PicturePath = tempCourse.PicturePath
		userMaxCourseDto.IsActive = tempCourse.IsActive

		// Declaramos la variable para llenar las categorias
		var categoriesDto dto.CategoriesMaxDto

		// Conseguimos las categorias de ese curso.
		tempCourseCategories, err := courseClient.GetCategoriesByCourseId(userCourse.IdCourse)

		if err != nil {
			return nil, err
		}

		for _, category := range tempCourseCategories { // Recorre las categorías de cada curso
			var categoryDto dto.CategoryMaxDto

			categoryDto.IdCategory = category.IdCategory
			categoryDto.Name = category.Name

			categoriesDto = append(categoriesDto, categoryDto)
		}

		userMaxCourseDto.Categories = categoriesDto
		userCoursesDto = append(userCoursesDto, userMaxCourseDto)
	}

	return userCoursesDto, nil
}

func (s *usersService) CheckAdmin(token string) (bool, e.ApiError) {
	id, err := s.ValidateToken(token)

	if err != nil {
		return false, err
	}

	user, err := s.GetUserById(id)

	if err != nil {
		return false, err
	}

	return user.IsAdmin, nil
}

// Remove user from usercourse
func (s *usersService) UnsubscribeUserCourse(courseId int, token string) e.ApiError {
	idUser, err := s.ValidateToken(token)

	if err != nil {
		return err
	}

	err1 := usersClient.RemoveUserCourse(idUser, courseId)

	if err1 != nil {
		return err1
	}

	return nil
}
