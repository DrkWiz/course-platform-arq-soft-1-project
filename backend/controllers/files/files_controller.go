package files

import (
	s "backend/services/files"

	e "backend/utils/errors"
	"io"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetFileById(c *gin.Context) { // Con el id del archivo en la BD devuelve la info guardaada en la BD.
	id := c.Param("id")
	idFile, err := strconv.Atoi(id)
	if err != nil {
		log.Info("Error parsing id")
		c.JSON(http.StatusBadRequest, e.NewBadRequestApiError("Error parsing id"))
		return
	}

	fileMinDto, err1 := s.FileService.GetFileById(idFile)

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusOK, fileMinDto)
}

func UploadFile(c *gin.Context) {
	// Check Authorization header for token
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, e.NewUnauthorizedApiError("Authorization header is required"))
		return
	}
	tokenParts := strings.Split(authHeader, "Bearer ")
	if len(tokenParts) < 2 || tokenParts[1] == "" {
		c.JSON(http.StatusUnauthorized, e.NewUnauthorizedApiError("Invalid token format"))
		return
	}
	token := tokenParts[1]

	// Parse course ID
	idCourse, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Info("Error parsing course ID")
		c.JSON(http.StatusBadRequest, e.NewBadRequestApiError("Error parsing course ID"))
		return
	}

	// Check file existence in form data
	file, err := c.FormFile("file")
	if err != nil {
		log.Info("Error getting file")
		c.JSON(http.StatusBadRequest, e.NewBadRequestApiError("Error getting file"))
		return
	}

	// Open the uploaded file
	openFile, err := file.Open()
	if err != nil {
		log.Info("Error opening file")
		c.JSON(http.StatusBadRequest, e.NewBadRequestApiError("Error opening file"))
		return
	}
	defer openFile.Close()

	// Read file content
	fileBytes, err := io.ReadAll(openFile)
	if err != nil {
		log.Info("Error reading file")
		c.JSON(http.StatusBadRequest, e.NewBadRequestApiError("Error reading file"))
		return
	}

	// Define the file path to save
	path := filepath.Join("./uploads/files/", file.Filename)

	// Call service to save the file
	err = s.FileService.SaveFile(fileBytes, path, idCourse, token)
	if err != nil {
		c.JSON(e.NewInternalServerApiError("Could not save file", err).Status(), err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"file_path": path})
}

func GetFilesByCourse(c *gin.Context) { // Con el id del curso devuelve un array de files min dto con los paths.
	idCourse, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Info("Error parsing course ID")
		c.JSON(http.StatusBadRequest, e.NewBadRequestApiError("Error parsing course ID"))
		return
	}

	files, err1 := s.FileService.GetFilesByCourse(idCourse)

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusOK, files)
}

func DownloadFile(c *gin.Context) { // Con el id del archivo en la BD devuelve el archivo para descargar.
	idFile, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Info("Error parsing id")
		c.JSON(http.StatusBadRequest, e.NewBadRequestApiError("Error parsing id"))
		return
	}

	fileMinDto, err1 := s.FileService.GetFileById(idFile)

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.FileAttachment(fileMinDto.Path, fileMinDto.Name)
	c.JSON(http.StatusOK, fileMinDto)
}
