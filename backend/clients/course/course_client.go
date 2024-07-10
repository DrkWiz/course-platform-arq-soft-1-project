package course

import (
	categoryModel "backend/model/category"
	courseModel "backend/model/courses"
	userCourseModel "backend/model/users"
	"os"

	e "backend/utils/errors"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB // aca se guarda la conexion a la base de datos

func GetCourseById(id int) (courseModel.Course, e.ApiError) { // aca se obtiene un curso por id
	var course courseModel.Course // se crea una variable course de tipo Course

	err := Db.Where("id_course = ?", id).First(&course).Error // se busca el curso por id

	if err != nil {
		return course, e.NewNotFoundApiError("Course not found")
	}

	log.Debug("Course: ", course)

	return course, nil
}

func CreateCourse(course *courseModel.Course) e.ApiError { // aca se crea un curso
	err := Db.Create(&course).Error // se crea el curso
	if err != nil {
		return e.NewInternalServerApiError("Error creating course", err) // si hay un error se retorna un error
	}
	return nil

}

func UpdateCourse(course courseModel.Course) e.ApiError { // aca se actualiza un curso
	err := Db.Save(&course).Error // save es un metodo de gorm que actualiza un registro

	if err != nil {
		return e.NewInternalServerApiError("Error updating course", err)
	}
	return nil
}

func DeleteCourse(id int) e.ApiError { // aca se elimina un curso por id
	err := Db.Model(&courseModel.Course{}).Where("id_course = ?", id).Update("is_active", "0").Error // se actualiza el campo is_active a 0
	if err != nil {
		return e.NewInternalServerApiError("Error deleting course", err)
	}
	return nil
}

func CreateCourseCategory(courseId int, categoryId int) e.ApiError { // aca se crea una categoria de un curso
	courseCategory := courseModel.CourseCategory{IdCourse: courseId, IdCategory: categoryId} // se crea una variable courseCategory de tipo CourseCategory
	err := Db.Create(&courseCategory).Error                                                  // se crea la categoria del curso en la base de datos
	if err != nil {
		return e.NewInternalServerApiError("Error creating course category", err)
	}
	return nil
}

func GetCategoriesByCourseId(id int) (categoryModel.Categories, e.ApiError) { // aca se obtienen las categorias de un curso por id
	var categories []categoryModel.Category                                                                                                                    // se crea una variable categories de tipo Category
	err := Db.Raw("SELECT * FROM categories WHERE id_category IN (SELECT id_category FROM course_categories WHERE id_course = ?)", id).Scan(&categories).Error // se hace una consulta a la base de datos para obtener las categorias de un curso

	if err != nil {
		return nil, e.NewNotFoundApiError("Categories not found")
	}
	return categories, nil
}

func GetCourses() (courseModel.Courses, e.ApiError) { // aca se obtienen todos los cursos
	var courses []courseModel.Course // se crea una variable courses de tipo Course que es un arreglo de cursos
	err := Db.Find(&courses).Error   // se buscan todos los cursos en la base de datos

	if err != nil {
		return nil, e.NewNotFoundApiError("Courses not found")
	}

	return courses, nil
}

func GetOwner(courseId int) (int, e.ApiError) { // aca se obtiene el dueño de un curso por id
	var course courseModel.Course                                   // se crea una variable course de tipo Course
	err := Db.Where("id_course = ?", courseId).First(&course).Error // se busca el curso por id
	if err != nil {
		return 0, e.NewNotFoundApiError("Course not found")
	}
	return course.IdOwner, nil
}

func SaveFile(file []byte, path string) e.ApiError { // aca se guarda un archivo
	err := os.WriteFile(path, file, 0644) // se guarda el archivo en la ruta especificada con los permisos 0644 (lectura y escritura)
	// 0644 es un permiso de lectura y escritura para el propietario y solo lectura para los demás usuarios del sistema operativo
	if err != nil {
		return e.NewInternalServerApiError("Error saving file", err)
	}
	return nil
}

func GetFile(path string) ([]byte, e.ApiError) {
	file, err := os.ReadFile(path) // se lee el archivo de la ruta especificada
	if err != nil {
		return nil, e.NewInternalServerApiError("Error getting file", err)
	}
	return file, nil // se retorna el archivo leido en bytes
}

func GetAvgRating(courseId int) (float64, e.ApiError) { // se busca un promedio de calificacion de un curso
	var userCourse userCourseModel.UserCourses                                                                                            // se crea una variable userCourse de tipo UserCourse
	err := Db.Raw("SELECT AVG(rating) as rating FROM user_courses WHERE id_course = ? AND rating <> 0", courseId).Scan(&userCourse).Error // se hace una consulta a la base de datos para obtener el promedio de calificacion de un curso
	if err != nil {
		return 0, e.NewNotFoundApiError("Rating not found")
	}
	return userCourse.Rating, nil
}

func GetComments(courseId int) ([]userCourseModel.UserCourses, e.ApiError) { // se obtienen los comentarios de un curso
	var userCourse []userCourseModel.UserCourses                                                                        // se crea una variable userCourses de tipo UserCourses
	err := Db.Raw("SELECT * FROM user_courses WHERE id_course = ? AND comment <> ''", courseId).Scan(&userCourse).Error // se hace una consulta a la base de datos para obtener los comentarios de un curso
	if err != nil {
		return nil, e.NewNotFoundApiError("Comments not found")
	}
	return userCourse, nil
}

func SetComment(courseId int, userId int, comment string) e.ApiError { // se crea un comentario actualizando el campo comment de la tabla user_courses
	err := Db.Model(&userCourseModel.UserCourses{}).Where("id_course = ? AND id_user = ?", courseId, userId).Update("comment", comment).Error // se actualiza el campo comment de la tabla user_courses
	if err != nil {
		return e.NewInternalServerApiError("Error setting comment", err)
	}
	return nil
}

func SetRating(courseId int, userId int, rating float32) e.ApiError { // Actualizamos la calificacion en la tabla user courses
	err := Db.Model(&userCourseModel.UserCourses{}).Where("id_course = ? AND id_user = ?", courseId, userId).Update("rating", rating).Error // se actualiza el campo rating de la tabla user_courses
	if err != nil {
		return e.NewInternalServerApiError("Error setting rating", err)
	}
	return nil
}
