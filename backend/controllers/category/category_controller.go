package category

import (
	"backend/dto"
	s "backend/services/category"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetCategoryById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid ID")
		return
	}

	log.Print("GetCategoryByUd: ", id)

	response, errUno := s.CategoryService.GetCategoryById(id)

	if errUno != nil {
		c.JSON(errUno.Status(), errUno)
		return
	}

	c.JSON(http.StatusOK, response)
}

func CreateCategory(c *gin.Context) {
	var category dto.CategoryCreateDto

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s.CategoryService.CreateCategory(category)
	c.JSON(http.StatusOK, "Category created")
}

func GetCategories(c *gin.Context) {
	log.Print("GetCategories")

	response, err := s.CategoryService.GetCategories()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, response)
}
