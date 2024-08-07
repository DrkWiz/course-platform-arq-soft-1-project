package courses

import (
	courseClient "backend/clients/course"
	"backend/dto"
	usersService "backend/services/users"

	courseModel "backend/model/courses"

	e "backend/utils/errors"

	log "github.com/sirupsen/logrus"
)

type coursesService struct{} // aca se define la clase CoursesService

type coursesServiceInterface interface { // aca se define la interfaz de la clase CoursesService con los metodos que se van a implementar
	GetCourseById(id int) (dto.CourseMaxDto, e.ApiError)
	CreateCourse(course dto.CourseCreateDto, token string) e.ApiError
	UpdateCourse(id int, course dto.CourseUpdateDto, token string) e.ApiError
	DeleteCourse(id int) e.ApiError
	GetCourses() (dto.CoursesMaxDto, e.ApiError)
	CheckOwner(token string, courseId int) (bool, e.ApiError)
	SaveFile(file []byte, path string) e.ApiError
	GetFile(path string) ([]byte, e.ApiError)
	GetAvgRating(courseId int) (float64, e.ApiError)
	GetComments(courseId int) ([]dto.CommentMaxDto, e.ApiError)
	SetComment(courseId int, userId int, comment string) e.ApiError
	SetRating(courseId int, userId int, rating float64) e.ApiError
}

var ( // aca se crea una variable CoursesService de tipo coursesServiceInterface
	CoursesService coursesServiceInterface
)

func init() { // aca se inicializa la variable CoursesService
	CoursesService = &coursesService{}
}

func (s *coursesService) GetCourseById(id int) (dto.CourseMaxDto, e.ApiError) { // aca se implementa el metodo GetCourseById de la interfaz CoursesServiceInterface
	log.Print("GetCourseById: ", id)

	course, err := courseClient.GetCourseById(id) // aca se llama al metodo GetCourseById del cliente CourseClient
	if err != nil {
		return dto.CourseMaxDto{}, err
	}

	// Create the CourseMaxDto
	courseMaxDto := dto.CourseMaxDto{ // aca se crea una variable de tipo CourseMaxDto
		IdCourse:    course.IdCourse,
		Owner:       course.IdOwner,
		Name:        course.Name,
		Description: course.Description,
		Price:       course.Price,
		PicturePath: course.PicturePath,
		StartDate:   course.StartDate,
		EndDate:     course.EndDate,
		IsActive:    course.IsActive,
		Categories:  []dto.CategoryMaxDto{},
	}

	// Fetch categories for the course
	tempCourses, err := courseClient.GetCategoriesByCourseId(course.IdCourse) // aca se llama al metodo GetCategoriesByCourseId del cliente CourseClient
	if err != nil {
		return dto.CourseMaxDto{}, err
	}

	// Append categories to the CourseMaxDto
	for _, category := range tempCourses { // aca se recorre la lista de categorias y se las agrega a la variable CourseMaxDto
		courseMaxDto.Categories = append(courseMaxDto.Categories, dto.CategoryMaxDto{
			IdCategory: category.IdCategory,
			Name:       category.Name,
		})
	}

	log.Println("CourseMaxDto: ", courseMaxDto)

	return courseMaxDto, nil // aca se devuelve la variable CourseMaxDto
}

// Create course
func (s *coursesService) CreateCourse(course dto.CourseCreateDto, token string) e.ApiError { // aca se implementa el metodo CreateCourse de la interfaz CoursesServiceInterface

	ownerId, err := usersService.UsersService.ValidateToken(token) // aca se llama al metodo ValidateToken del servicio UsersService para obtener el id del usuario

	if err != nil {
		return err
	}

	ownerDTO, err := usersService.UsersService.GetUserById(ownerId) // aca se llama al metodo GetUserById del servicio UsersService para obtener el usuario

	if err != nil {
		return err
	}

	isAdmin := ownerDTO.IsAdmin // aca se obtiene si el usuario es admin
	if !isAdmin {               // aca se llama al metodo GetUserById del cliente UsersClient y se verifica si el usuario es admin
		return e.NewForbiddenApiError("You don't have permission to create a course")
	}

	courseToCreate := courseModel.Course{
		Name:        course.Name,
		Description: course.Description,
		Price:       course.Price,
		PicturePath: course.PicturePath,
		StartDate:   course.StartDate,
		EndDate:     course.EndDate,
		IdOwner:     ownerId} // aca se crea una variable de tipo Course con los datos del curso a crear

	err = courseClient.CreateCourse(&courseToCreate) // aca se llama al metodo CreateCourse del cliente CourseClient y se le pasa la variable Course a crear como parametro para crear el curso

	if err != nil {
		return err
	}

	log.Println("Course created: ", courseToCreate)

	for _, categoryId := range course.CategoriesId { // aca se recorre la lista de categorias y se las agrega al curso creado con el metodo CreateCourseCategory del cliente CourseClient
		err = courseClient.CreateCourseCategory(courseToCreate.IdCourse, categoryId)
		if err != nil {
			return err
		}
	}

	return nil

}

// Update a course.

func (s *coursesService) UpdateCourse(courseId int, course dto.CourseUpdateDto, token string) e.ApiError { // aca se implementa el metodo UpdateCourse de la interfaz CoursesServiceInterface

	ownerId, err := courseClient.GetOwner(courseId)

	if err != nil {
		return err
	}

	ok, err := usersService.UsersService.CheckAdmin(token)

	if err != nil {
		return err
	}

	if !ok { // aca se verifica si el usuario es admin para poder actualizar el curso
		isOwner, err := s.CheckOwner(token, courseId) // aca se llama al metodo CheckOwner del servicio CoursesService para verificar si el usuario es el dueño del curso
		if err != nil {
			return err
		}
		if !isOwner { // aca se verifica si el usuario es el dueño del curso
			return e.NewForbiddenApiError("You don't have permission to update this course") // aca se devuelve un error si el usuario no es el dueño del curso
		}
	}

	originalCourse, err := courseClient.GetCourseById(courseId) // Obtenemos el curso original

	if err != nil {
		return err
	}

	courseToUpdate := courseModel.Course{ // aca se crea una variable de tipo Course con los datos del curso a actualizar
		IdCourse:    courseId,
		Name:        course.Name,
		Description: course.Description,
		Price:       course.Price,
		PicturePath: course.PicturePath,
		StartDate:   course.StartDate,
		EndDate:     course.EndDate,
		IdOwner:     ownerId,
		IsActive:    course.IsActive}

	err = courseClient.UpdateCourse(courseToUpdate) // aca se llama al metodo UpdateCourse del cliente CourseClient y se le pasa la variable Course a actualizar como parametro para actualizar el curso

	log.Println("Course updated: ", courseToUpdate)

	if err != nil {
		log.Println("Error updating course: ", err.Status())
		err := courseClient.UpdateCourse(originalCourse) //  Rollback
		if err != nil {
			return e.NewInternalServerApiError("Error restoring course", err)
		}
		return err
	}

	// Inicia la actualizacion de la tabla puente course_category

	originalCategoriesIds, err := courseClient.GetCategoriesByCourseId(courseId) // aca se llama al metodo GetCategoriesByCourseId del cliente CourseClient para obtener las categorias del curso
	log.Println("Original categories: ", originalCategoriesIds)

	if err != nil {
		return err
	}

	categoriesIds := []int{}

	for _, category := range course.Categories { // aca se recorre la lista de categorias y se las agrega a la variable categoriesIds
		categoriesIds = append(categoriesIds, category.IdCategory)
	}
	log.Println("New categories: ", categoriesIds)

	temp := 0

	for _, category := range originalCategoriesIds { // aca se recorre la lista de categorias originales y se las elimina con el metodo DeleteCourseCategory del cliente CourseClient
		err = courseClient.DeleteCourseCategory(courseId, category.IdCategory)
		temp = temp + 1
		if err != nil {
			break
		}
	}

	if err != nil {
		for i := 0; i < temp; i++ {
			err = courseClient.CreateCourseCategory(courseId, originalCategoriesIds[i].IdCategory)
			if err != nil {
				return e.NewInternalServerApiError("Error restoring course categories", err)
			}
		}
		return err
	}
	log.Println("Old Categories deleted")

	temp = 0

	for _, newCategory := range categoriesIds { // aca se recorre la lista de categorias nuevas y se las agrega con el metodo CreateCourseCategory del cliente CourseClient
		err = courseClient.CreateCourseCategory(courseId, newCategory)
		temp = temp + 1
		if err != nil {
			break
		}
	}

	if err != nil {
		for i := 0; i < temp; i++ {
			err = courseClient.DeleteCourseCategory(courseId, categoriesIds[i])
			if err != nil {
				return e.NewInternalServerApiError("Error restoring course categories", err)
			}
		}
		for _, category := range originalCategoriesIds {
			err = courseClient.CreateCourseCategory(courseId, category.IdCategory)
			if err != nil {
				return e.NewInternalServerApiError("Error restoring course categories", err)
			}
		}
		return err
	}

	log.Println("Course updated: ", courseToUpdate) // aca se imprime el curso actualizado

	return nil
}

//Soft delete course
func (s *coursesService) DeleteCourse(id int) e.ApiError { // aca se implementa el metodo DeleteCourse de la interfaz CoursesServiceInterface
	err := courseClient.DeleteCourse(id) // aca se llama al metodo DeleteCourse del cliente CourseClient para eliminar el curso
	if err != nil {
		return err
	}

	return nil
}

// Get all courses in DB
func (s *coursesService) GetCourses() (dto.CoursesMaxDto, e.ApiError) { // aca se implementa el metodo GetCourses de la interfaz CoursesServiceInterface

	courses, err := courseClient.GetCourses() // aca se llama al metodo GetCourses del cliente CourseClient para obtener todos los cursos

	if err != nil {
		return nil, err
	}

	var CoursesMaxDto dto.CoursesMaxDto // aca se crea una variable de tipo CoursesMaxDto que es una lista de CourseMaxDto

	for _, course := range courses { // aca se recorre la lista de cursos y se los agrega a la variable CoursesMaxDto
		CourseMaxDto := dto.CourseMaxDto{ // aca se crea una variable de tipo CourseMaxDto
			IdCourse:    course.IdCourse,
			Owner:       course.IdOwner,
			Name:        course.Name,
			Description: course.Description,
			Price:       course.Price,
			PicturePath: course.PicturePath,
			StartDate:   course.StartDate,
			EndDate:     course.EndDate,
			IsActive:    course.IsActive}
		tempCourses, err := courseClient.GetCategoriesByCourseId(course.IdCourse) // aca se llama al metodo GetCategoriesByCourseId del cliente CourseClient para obtener las categorias de cada curso

		if err != nil {
			return nil, err
		}

		for _, category := range tempCourses { // aca se recorre la lista de categorias y se las agrega a la variable CourseMaxDto
			CourseMaxDto.Categories = append(CourseMaxDto.Categories, dto.CategoryMaxDto{IdCategory: category.IdCategory, Name: category.Name}) // aca se agrega la categoria a la variable CourseMaxDto creada , append agrega un elemento al final de la lista
		}
		CoursesMaxDto = append(CoursesMaxDto, CourseMaxDto) // aca se agrega el curso a la variable CoursesMaxDto creada, append agrega un elemento al final de la lista
	}

	return CoursesMaxDto, nil
}

// Check if token is the owner of the course
func (s *coursesService) CheckOwner(token string, courseId int) (bool, e.ApiError) {
	idToCheck, err := usersService.UsersService.ValidateToken(token) // aca se llama al metodo ValidateToken del servicio UsersService para obtener el id del usuario

	if err != nil {
		return false, err
	}

	ownerId, err := courseClient.GetOwner(courseId) // aca se llama al metodo GetOwner del cliente CourseClient para obtener el id del dueño del curso

	if err != nil {
		return false, err
	}

	return ownerId == idToCheck, nil // aca se devuelve true si el id del dueño del curso es igual al id del usuario
}

func (s *coursesService) SaveFile(file []byte, path string) e.ApiError { // aca se implementa el metodo SaveFile de la interfaz CoursesServiceInterface
	err := courseClient.SaveFile(file, path) // aca se llama al metodo SaveFile del cliente CourseClient para guardar el archivo
	if err != nil {
		return err
	}
	return nil
}

func (s *coursesService) GetFile(path string) ([]byte, e.ApiError) { // aca se implementa el metodo GetFile de la interfaz CoursesServiceInterface
	file, err := courseClient.GetFile(path) // aca se llama al metodo GetFile del cliente CourseClient para obtener el archivo
	if err != nil {
		return nil, err
	}
	return file, nil
}

// avg rating
func (s *coursesService) GetAvgRating(courseId int) (float64, e.ApiError) { // aca se implementa el metodo GetAvgRating de la interfaz CoursesServiceInterface
	rating, err := courseClient.GetAvgRating(courseId) // aca se llama al metodo GetAvgRating del cliente CourseClient para obtener el rating promedio
	if err != nil {
		return 0, err
	}
	return rating, nil
}

// comments
func (s *coursesService) GetComments(courseId int) ([]dto.CommentMaxDto, e.ApiError) { // aca se implementa el metodo GetComments de la interfaz CoursesServiceInterface
	comments, err := courseClient.GetComments(courseId) // aca se llama al metodo GetComments del cliente CourseClient para obtener los comentarios
	if err != nil {
		return nil, err
	}
	var CommentsMaxDto []dto.CommentMaxDto // aca se crea una variable de tipo CommentMaxDto que es una lista de comentarios
	for _, comment := range comments {     // aca se recorre la lista de comentarios y se los agrega a la variable CommentsMaxDto

		userDTO, err := usersService.UsersService.GetUserById(comment.IdUser) // aca se llama al metodo GetUserById del servicio UsersService para obtener el usuario del comentario
		if err != nil {
			return nil, err
		}

		CommentMaxDto := dto.CommentMaxDto{ // aca se crea una variable de tipo CommentMaxDto
			Username: userDTO.Username, // aca se llama al metodo GetUserById del cliente UsersClient para obtener el nombre de usuario del comentario
			Comment:  comment.Comment,
			Rating:   comment.Rating}

		CommentsMaxDto = append(CommentsMaxDto, CommentMaxDto) // aca se agrega el comentario a la variable CommentsMaxDto creada, append agrega un elemento al final de la lista

	}
	return CommentsMaxDto, nil
}

// set comment
func (s *coursesService) SetComment(courseId int, userId int, comment string) e.ApiError { // aca se implementa el metodo SetComment de la interfaz CoursesServiceInterface
	err := courseClient.SetComment(courseId, userId, comment) // aca se llama al metodo SetComment del cliente CourseClient para crear un comentario
	if err != nil {
		return err
	}
	return nil
}

// set rating
func (s *coursesService) SetRating(courseId int, userId int, rating float64) e.ApiError { // aca se implementa el metodo SetRating de la interfaz CoursesServiceInterface
	err := courseClient.SetRating(courseId, userId, rating) // aca se llama al metodo SetRating del cliente CourseClient para crear un rating
	if err != nil {
		return err
	}
	return nil
}
