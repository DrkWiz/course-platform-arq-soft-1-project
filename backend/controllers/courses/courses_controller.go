package courses

import (
	"backend/dto"
	s "backend/services/courses"
	"io"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	e "backend/utils/errors"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	usersService "backend/services/users"
)

// Get Course by ID
func GetCourseById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id")) //aca se obtiene el id del curso que se quiere buscar, y se lo convierte  de str a int

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid ID")
		return
	}

	log.Print("GetCourseById: ", id)

	response, err1 := s.CoursesService.GetCourseById(id) // aca se llama a la funcion GetCourseById de CoursesService

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusOK, response) //aca se devuelve el curso encontrado
}

// Create new course
func CreateCourse(c *gin.Context) {
	authHeader := c.GetHeader("Authorization") //aca se obtiene el token del header de la peticion

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, e.NewUnauthorizedApiError("Authorization header is required"))
		return
	}

	token := strings.Split(authHeader, "Bearer ")[1] //aca se obtiene el token del header y se lo separa para obtener solo el token en si mismo (sin el Bearer)

	if token == "" {
		c.JSON(http.StatusUnauthorized, e.NewUnauthorizedApiError("Token is required"))
		return
	}

	var course dto.CourseCreateDto //aca se crea una variable de tipo CourseCreateDto
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) //aca se devuelve un error si no se pudo hacer el binding del JSON
		return
	}

	s.CoursesService.CreateCourse(course, token) //aca se llama a la funcion CreateCourse de la interfaz CoursesService y se le pasa el curso y el token
	c.JSON(http.StatusCreated, "Course created") //aca se devuelve un mensaje de que el curso fue creado exitosamente
}

// Update course
func UpdateCourse(c *gin.Context) {
	authHeader := c.GetHeader("Authorization") //aca se obtiene el token del header de la peticion

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, e.NewUnauthorizedApiError("Authorization header is required"))
		return
	}

	token := strings.Split(authHeader, "Bearer ")[1] //aca se obtiene el token del header y se lo separa para obtener solo el token en si mismo (sin el Bearer)
	if token == "" {
		c.JSON(http.StatusUnauthorized, e.NewUnauthorizedApiError("Token is required"))
		return
	}

	var course dto.CourseUpdateDto //aca se crea una variable de tipo CourseUpdateDto
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, e.NewBadRequestApiError("Invalid JSON body"))
		return
	}

	id, err := strconv.Atoi(c.Param("id")) //aca se obtiene el id del curso que se quiere actualizar, y se lo convierte  de str a int

	if err != nil {
		c.JSON(http.StatusBadRequest, "Bad ID")
		return
	}

	err1 := s.CoursesService.UpdateCourse(id, course, token) //aca se llama a la funcion UpdateCourse de la interfaz CoursesService
	//y se le pasa el id del curso, el curso y el token

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusNoContent, "Course updated") //aca se devuelve un mensaje de que el curso fue actualizado exitosamente

}

// Soft delete course
func DeleteCourse(c *gin.Context) {
	authHeader := c.GetHeader("Authorization") //aca se obtiene el token del header de la peticion

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, e.NewUnauthorizedApiError("Authorization header is required"))
		return
	}

	token := strings.Split(authHeader, "Bearer ")[1] //el [1] es para obtener solo el token en si mismo (sin el Bearer)
	if token == "" {
		c.JSON(http.StatusUnauthorized, e.NewUnauthorizedApiError("Token is required"))
		return
	}

	id, err := strconv.Atoi(c.Param("id")) //aca se obtiene el id del curso que se quiere borrar, y se lo convierte  de str a int

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid ID")
		return
	}

	err1 := s.CoursesService.DeleteCourse(id) //aca se llama a la funcion DeleteCourse de la interfaz CoursesService y se le pasa el id del curso

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusNoContent, "Course deleted") //aca se devuelve un mensaje de que el curso fue borrado exitosamente
}

// Get all courses in db
func GetCourses(c *gin.Context) {
	response, err1 := s.CoursesService.GetCourses() //aca se llama a la funcion GetCourses de la interfaz CoursesService

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusOK, response) //aca se devuelven todos los cursos encontrados
}

// Check if the token is the owner of the course
func CheckOwner(c *gin.Context) {
	authHeader := c.GetHeader("Authorization") //aca se obtiene el token del header de la peticion

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, e.NewUnauthorizedApiError("Authorization header is required"))
		return
	}

	token := strings.Split(authHeader, "Bearer ")[1] //aca se obtiene el token del header y se lo separa para obtener solo el token en si mismo (sin el Bearer)
	if token == "" {
		c.JSON(http.StatusUnauthorized, e.NewUnauthorizedApiError("Token is required"))
		return
	}

	courseId, err := strconv.Atoi(c.Param("id")) //aca se obtiene el id del curso que se quiere verificar, y se lo convierte  de str a int

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid ID")
		return
	}

	response, err1 := s.CoursesService.CheckOwner(token, courseId) //aca se llama a la funcion CheckOwner de la interfaz CoursesService

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusOK, response) //aca se devuelve si el token es el dueño del curso o no
}

// Upload image
func ImageUpload(c *gin.Context) {
	file, err := c.FormFile("image") //aca se obtiene el archivo de la peticion

	if err != nil {
		c.JSON(http.StatusBadRequest, e.NewBadRequestApiError("Error getting file"))
		return
	}

	path := filepath.Join("./uploads", file.Filename) //aca se obtiene el path donde se va a guardar el archivo

	openFile, err := file.Open() //aca se abre el archivo para leerlo y guardarlo en la variable openFile para corroborar que no haya errores
	if err != nil {
		c.JSON(http.StatusInternalServerError, e.NewInternalServerApiError("Error opening file", err))
		return
	}
	defer openFile.Close() //aca se cierra el archivo al finalizar la funcion

	fileBytes, err := io.ReadAll(openFile) //aca se lee el archivo y se guarda en la variable fileBytes	de tipo []byte

	if err != nil {
		c.JSON(http.StatusInternalServerError, e.NewInternalServerApiError("Error reading file", err))
		return
	}

	err = s.CoursesService.SaveFile(fileBytes, path) //aca se llama a la funcion SaveFile de la interfaz CoursesService

	c.JSON(http.StatusCreated, gin.H{"picture_path": file.Filename}) //aca se devuelve el path relativo de la imagen guardada en el servidor	 para ser usado en el frontend
}

// Devuelve el path de la imagen
func GetImage(c *gin.Context) {
	picturepath := c.Param("picturepath")           //aca se obtiene el path de la imagen que se quiere obtener
	path := filepath.Join("./uploads", picturepath) //aca se obtiene el path completo de la imagen

	file, err := s.CoursesService.GetFile(path) //aca se llama a la funcion GetFile de la interfaz CoursesService

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.Data(http.StatusOK, "image/jpeg", file) //aca se devuelve la imagen en formato jpeg
}

// Devuelve el rating promedio de un curso.
func GetAvgRating(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id")) //aca se obtiene el id del curso que se quiere buscar, y se lo convierte  de str a int

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid ID")
		return
	}

	response, err1 := s.CoursesService.GetAvgRating(id) // aca se llama a la funcion GetAvgRating de la interfaz CoursesService

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusOK, response) //aca se devuelve el promedio de calificacion del curso
}

// Devuelve los comentarios de un curso.
func GetComments(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id")) //aca se obtiene el id del curso que se quiere buscar, y se lo convierte  de str a int

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid ID")
		return
	}

	response, err1 := s.CoursesService.GetComments(id) // aca se llama a la funcion GetComments de la interfaz CoursesService

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}
	log.Print("GetComments: ", response)
	c.JSON(http.StatusOK, response) //aca se devuelven los comentarios del curso
}

// Actualiza un comentario.
func SetComment(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")

	if !strings.HasPrefix(tokenStr, "Bearer ") {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

	courseId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Invalid course ID: %v", err)
		c.JSON(http.StatusBadRequest, "Invalid ID")
		return
	}

	userId, err := usersService.UsersService.ValidateToken(tokenStr)
	if err != nil {
		log.Printf("Invalid token: %v", err)
		c.JSON(http.StatusUnauthorized, "Invalid token")
		return
	}

	var comment dto.CommentCreateDto
	if err := c.ShouldBindJSON(&comment); err != nil {
		log.Printf("Invalid JSON body: %v", err)
		c.JSON(http.StatusBadRequest, e.NewBadRequestApiError("Invalid JSON body"))
		return
	}

	log.Printf("Adding comment: courseId=%d, userId=%d, comment=%s", courseId, userId, comment.Comment)
	err = s.CoursesService.SetComment(courseId, userId, comment.Comment)
	if err != nil {
		log.Printf("Error adding comment: %v", err)
		c.JSON(http.StatusBadRequest, e.NewBadRequestApiError("Comment not added"))
		return
	}

	c.JSON(http.StatusNoContent, "Comment added")
}

// Actualiza un rating
func SetRating(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")

	if !strings.HasPrefix(tokenStr, "Bearer ") {
		c.JSON(http.StatusUnauthorized, "Token is required")
		return
	}

	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

	if tokenStr == "" {
		c.JSON(http.StatusUnauthorized, "Token is required")
		return
	}

	courseId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Invalid course ID: %v", err)
		c.JSON(http.StatusBadRequest, "Invalid ID")
		return
	}

	userId, errApi := usersService.UsersService.ValidateToken(tokenStr)

	if errApi != nil {
		log.Printf("Invalid token: %v", errApi.Error())
		c.JSON(errApi.Status(), errApi)
		return
	}

	var requestBody struct {
		Rating float64 `json:"rating"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		log.Printf("Invalid JSON input: %v", err)
		c.JSON(http.StatusBadRequest, "Invalid input")
		return
	}

	rating := requestBody.Rating

	log.Printf("Adding rating: courseId=%d, userId=%d, rating=%f", courseId, userId, rating)
	errApi = s.CoursesService.SetRating(courseId, userId, rating)
	if errApi != nil {
		log.Printf("Error adding rating: %v", errApi.Error())
		c.JSON(errApi.Status(), errApi)
		return
	}

	c.JSON(http.StatusNoContent, "Rating changed")
}

