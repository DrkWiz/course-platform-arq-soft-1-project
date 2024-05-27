package courses

import (
	categoryClient "backend/clients/category"
	courseClient "backend/clients/course"

	"backend/dto"

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
	CourseMinDto.Picture_path = course.Picture_path
	CourseMinDto.Start_date = course.Start_date
	CourseMinDto.End_date = course.End_date

	return CourseMinDto, nil
}

//Create course

func CreateCourse(course dto.CourseCreateDto) e.ApiError {

	courseToCreate := courseModel.Course{Name: course.Name, Description: course.Description, Price: course.Price, Picture_path: course.PicturePath, Start_date: course.StartDate, End_date: course.EndDate, Id_user: course.IdUser, IsActive: true}

	// Add course to db.
	id, err := courseClient.CreateCourse(courseToCreate)
	if err != nil {
		return err
	}

	// Add categories to course_categories in db.
	for _, idCategory := range course.IdCategories {
		//Check if category exists
		if categoryClient.CheckCategory(idCategory) {
			return e.NewBadRequestApiError("Category does not exist")
		}

		// Add the category.
		err := courseClient.AddCategoryToCourse(id, idCategory)
		if err != nil {
			return err
		}
	}

	return nil

}

// Update a course.

func UpdateCourse(id int, course dto.CourseUpdateDto) e.ApiError {
	courseToUpdate := courseModel.Course{IdCourse: id, Name: course.Name, Description: course.Description, Price: course.Price, Picture_path: course.Picture_path, Start_date: course.Start_date, End_date: course.End_date, Id_user: course.Id_user, IsActive: true}

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
