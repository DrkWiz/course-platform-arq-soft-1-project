package courses

import (
	courseClient "backend/clients/course"
	"backend/dto"
	usersService "backend/services/users"

	courseModel "backend/model/courses"

	e "backend/utils/errors"

	log "github.com/sirupsen/logrus"
)

type coursesService struct{}

type coursesServiceInterface interface {
	GetCourseById(id int) (dto.CourseMinDto, e.ApiError)
	CreateCourse(course dto.CourseCreateDto) e.ApiError
	UpdateCourse(id int, course dto.CourseUpdateDto, token string) e.ApiError
	DeleteCourse(id int) e.ApiError
	GetCourses() (dto.CoursesMaxDto, e.ApiError)
	CheckOwner(token string, courseId int) (bool, e.ApiError)
}

var (
	CoursesService coursesServiceInterface
)

func init() {
	CoursesService = &coursesService{}
}

func (s *coursesService) GetCourseById(id int) (dto.CourseMinDto, e.ApiError) {

	log.Print("GetCourseById: ", id)

	course, err := courseClient.GetCourseById(id)

	if err != nil {
		return dto.CourseMinDto{}, err
	}

	var CourseMinDto dto.CourseMinDto

	CourseMinDto.IdCourse = course.IdCourse
	CourseMinDto.Name = course.Name
	CourseMinDto.Description = course.Description
	CourseMinDto.Price = course.Price
	CourseMinDto.PicturePath = course.PicturePath
	CourseMinDto.StartDate = course.StartDate
	CourseMinDto.EndDate = course.EndDate
	CourseMinDto.IsActive = course.IsActive

	log.Println("CourseMinDto: ", CourseMinDto)

	return CourseMinDto, nil
}

//Create course

func (s *coursesService) CreateCourse(course dto.CourseCreateDto) e.ApiError {
	courseToCreate := courseModel.Course{Name: course.Name, Description: course.Description, Price: course.Price, PicturePath: course.PicturePath, StartDate: course.StartDate, EndDate: course.EndDate, IdOwner: course.IdOwner}
	
	err := courseClient.CreateCourse(courseToCreate)

	if err != nil {
		return err
	}
	
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

	courseToUpdate := courseModel.Course{IdCourse: courseId, Name: course.Name, Description: course.Description, Price: course.Price, PicturePath: course.PicturePath, StartDate: course.StartDate, EndDate: course.EndDate, IdOwner: course.IdOwner, IsActive: course.IsActive}

	err = courseClient.UpdateCourse(courseToUpdate)

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
