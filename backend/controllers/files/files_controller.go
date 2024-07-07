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

func UploadFile(c *gin.Context) { // Crea un archivo en la BD

	// Checkeamos Token.

	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, e.NewUnauthorizedApiError("Authorization header is required"))
		return
	}

	token := strings.Split(authHeader, "Bearer ")[1] //aca se obtiene el token del header y se lo separa para obtener solo el token en si mismo (sin el Bearer)
	if token == "" {
		c.JSON(http.StatusUnauthorized, e.NewUnauthorizedApiError("Token is required"))
		return
	}

	// Checkeamos ID.

	idCourse, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Info("Error parsing id")
		c.JSON(http.StatusBadRequest, e.NewBadRequestApiError("Error parsing id"))
		return
	}

	// Checkeamos el archivo.

	file, err := c.FormFile("file")
	if err != nil {
		log.Info("Error getting file")
		c.JSON(http.StatusBadRequest, e.NewBadRequestApiError("Error getting file"))
		return
	}

	path := filepath.Join("./uploads/files/", file.Filename)

	openFile, err := file.Open()
	if err != nil {
		log.Info("Error opening file")
		c.JSON(http.StatusBadRequest, e.NewBadRequestApiError("Error opening file"))
		return
	}
	defer openFile.Close()

	filebytes, err := io.ReadAll(openFile)
	if err != nil {
		log.Info("Error reading file")
		c.JSON(http.StatusBadRequest, e.NewBadRequestApiError("Error reading file"))
		return
	}

	// Llamamos al servicio para subir y crear el archivo.

	err1 := s.FileService.SaveFile(filebytes, path, idCourse, token)

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusOK, gin.H{"file_path": path})
}
