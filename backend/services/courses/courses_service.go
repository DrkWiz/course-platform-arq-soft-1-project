package courses

import (
	courseClient "backend/clients/course"
	usersClient "backend/clients/users"
	"backend/dto"
	usersService "backend/services/users"

	courseModel "backend/model/courses"

	e "backend/utils/errors"

	log "github.com/sirupsen/logrus"
)

type coursesService struct{}

type coursesServiceInterface interface {
	GetCourseById(id int) (dto.CourseMaxDto, e.ApiError)
	CreateCourse(course dto.CourseCreateDto, token string) e.ApiError
	UpdateCourse(id int, course dto.CourseUpdateDto, token string) e.ApiError
	DeleteCourse(id int) e.ApiError
	GetCourses() (dto.CoursesMaxDto, e.ApiError)
	CheckOwner(token string, courseId int) (bool, e.ApiError)
	SaveFile(file []byte, path string) e.ApiError
	GetFile(path string) ([]byte, e.ApiError)
}

var (
	CoursesService coursesServiceInterface
)

func init() {
	CoursesService = &coursesService{}
}

func (s *coursesService) GetCourseById(id int) (dto.CourseMaxDto, e.ApiError) {
	log.Print("GetCourseById: ", id)

	course, err := courseClient.GetCourseById(id)
	if err != nil {
		return dto.CourseMaxDto{}, err
	}

	// Create the CourseMaxDto
	courseMaxDto := dto.CourseMaxDto{
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
	tempCourses, err := courseClient.GetCategoriesByCourseId(course.IdCourse)
	if err != nil {
		return dto.CourseMaxDto{}, err
	}

	// Append categories to the CourseMaxDto
	for _, category := range tempCourses {
		courseMaxDto.Categories = append(courseMaxDto.Categories, dto.CategoryMaxDto{
			IdCategory: category.IdCategory,
			Name:       category.Name,
		})
	}

	log.Println("CourseMaxDto: ", courseMaxDto)

	return courseMaxDto, nil
}

//Create course

func (s *coursesService) CreateCourse(course dto.CourseCreateDto, token string) e.ApiError {

	ownerId, err := usersService.UsersService.ValidateToken(token)

	if err != nil {
		return err
	}

	if !usersClient.GetUserById(ownerId).IsAdmin {
		return e.NewForbiddenApiError("You don't have permission to create a course")
	}

	courseToCreate := courseModel.Course{Name: course.Name, Description: course.Description, Price: course.Price, PicturePath: course.PicturePath, StartDate: course.StartDate, EndDate: course.EndDate, IdOwner: ownerId}

	err = courseClient.CreateCourse(&courseToCreate)

	if err != nil {
		return err
	}

	log.Println("Course created: ", courseToCreate)

	for _, categoryId := range course.CategoriesId {
		err = courseClient.CreateCourseCategory(courseToCreate.IdCourse, categoryId)
		if err != nil {
			return err
		}
	}

	return nil

}

// Update a course.

func (s *coursesService) UpdateCourse(courseId int, course dto.CourseUpdateDto, token string) e.ApiError {

	ownerId, err := courseClient.GetOwner(courseId)

	if err != nil {
		return err
	}

	ok, err := usersService.UsersService.CheckAdmin(token)

	if err != nil {
		return err
	}

	if !ok {
		isOwner, err := s.CheckOwner(token, courseId)
		if err != nil {
			return err
		}
		if !isOwner {
			return e.NewForbiddenApiError("You don't have permission to update this course")
		}
	}

	courseToUpdate := courseModel.Course{IdCourse: courseId, Name: course.Name, Description: course.Description, Price: course.Price, PicturePath: course.PicturePath, StartDate: course.StartDate, EndDate: course.EndDate, IdOwner: ownerId, IsActive: course.IsActive}

	err = courseClient.UpdateCourse(courseToUpdate)

	log.Println("Course updated: ", courseToUpdate)

	if err != nil {
		return err
	}

	return nil
}

//Soft delete course

func (s *coursesService) DeleteCourse(id int) e.ApiError {
	err := courseClient.DeleteCourse(id)
	if err != nil {
		return err
	}

	return nil
}

// Get all courses in DB

func (s *coursesService) GetCourses() (dto.CoursesMaxDto, e.ApiError) {

	courses, err := courseClient.GetCourses()

	if err != nil {
		return nil, err
	}

	var CoursesMaxDto dto.CoursesMaxDto

	for _, course := range courses {
		CourseMaxDto := dto.CourseMaxDto{IdCourse: course.IdCourse, Owner: course.IdOwner, Name: course.Name, Description: course.Description, Price: course.Price, PicturePath: course.PicturePath, StartDate: course.StartDate, EndDate: course.EndDate, IsActive: course.IsActive}
		tempCourses, err := courseClient.GetCategoriesByCourseId(course.IdCourse)

		if err != nil {
			return nil, err
		}

		for _, category := range tempCourses {
			CourseMaxDto.Categories = append(CourseMaxDto.Categories, dto.CategoryMaxDto{IdCategory: category.IdCategory, Name: category.Name})
		}
		CoursesMaxDto = append(CoursesMaxDto, CourseMaxDto)
	}

	return CoursesMaxDto, nil
}

// Check if token is the owner of the course

func (s *coursesService) CheckOwner(token string, courseId int) (bool, e.ApiError) {
	idToCheck, err := usersService.UsersService.ValidateToken(token)

	if err != nil {
		return false, err
	}

	ownerId, err := courseClient.GetOwner(courseId)

	if err != nil {
		return false, err
	}

	return ownerId == idToCheck, nil
}

func (s *coursesService) SaveFile(file []byte, path string) e.ApiError {
	err := courseClient.SaveFile(file, path)
	if err != nil {
		return err
	}
	return nil
}

func (s *coursesService) GetFile(path string) ([]byte, e.ApiError) {
	file, err := courseClient.GetFile(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}
