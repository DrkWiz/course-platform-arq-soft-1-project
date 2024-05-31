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
}

var (
	CoursesService coursesServiceInterface
)

func init() {
	CoursesService = &coursesService{}
}

func (s *coursesService) GetCourseById(id int) (dto.CourseMinDto, e.ApiError) {

	log.Print("GetCourseById: ", id)

	var course courseModel.Course = courseClient.GetCourseById(id)
	var CourseMinDto dto.CourseMinDto

	CourseMinDto.IdCourse = course.IdCourse
	CourseMinDto.Name = course.Name
	CourseMinDto.Description = course.Description
	CourseMinDto.Price = course.Price
	CourseMinDto.PicturePath = course.PicturePath
	CourseMinDto.StartDate = course.Start_date
	CourseMinDto.EndDate = course.End_date

	return CourseMinDto, nil
}

//Create course

func CreateCourse(course dto.CourseCreateDto) error {

	courseToCreate := courseModel.Course{Name: course.Name, Description: course.Description, Price: course.Price, PicturePath: course.Picture_path, Start_date: course.Start_date, End_date: course.End_date, Id_user: course.Id_user}

	err := courseClient.CreateCourse(courseToCreate)
	if err != nil {
		return err
	}

	return nil

}

// Update a course.

func UpdateCourse(id int, course dto.CourseUpdateDto) e.ApiError {
	courseToUpdate := courseModel.Course{IdCourse: id, Name: course.Name, Description: course.Description, Price: course.Price, PicturePath: course.Picture_path, Start_date: course.Start_date, End_date: course.End_date, Id_user: course.Id_user, IsActive: true}

	err := courseClient.UpdateCourse(courseToUpdate)
	if err != nil {
		return err
	}

	return nil
}

//Soft delete course

func DeleteCourse(id int) error {

	err := courseClient.DeleteCourse(id)
	if err != nil {
		return err
	}

	return nil
}

// Get all courses in DB

func GetCourses() (dto.CoursesMaxDto, e.ApiError) {

	courses := courseClient.GetCourses()
	var CoursesMaxDto dto.CoursesMaxDto

	for _, course := range courses {
		CourseMaxDto := dto.CourseMaxDto{IdCourse: course.IdCourse, Name: course.Name, Description: course.Description, Price: course.Price, PicturePath: course.PicturePath, StartDate: course.Start_date, EndDate: course.End_date, IsActive: course.IsActive}
		tempCourses := courseClient.GetCategoriesByCourseId(course.IdCourse)
		for _, category := range tempCourses {
			CourseMaxDto.Categories = append(CourseMaxDto.Categories, dto.CategoryTokenDto{IdCategory: category.IdCategory, Name: category.Name})
		}
		CoursesMaxDto = append(CoursesMaxDto, CourseMaxDto)
	}

	return CoursesMaxDto, nil
}

// Check if token is the owner of the course

func CheckOwner(token string, courseId int) (bool, e.ApiError) {
	idToCheck, err := usersService.ValidateToken(token)

	if err != nil {
		return false, err
	}

	ownerId, err := courseClient.GetOwner(courseId)

	if err != nil {
		return false, err
	}

	return ownerId == idToCheck, nil
}
