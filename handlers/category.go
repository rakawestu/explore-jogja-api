package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rakawestu/explore-jogja-api/models"
	"github.com/rakawestu/explore-jogja-api/orm"
)

// GetAllCategoriesHandler returns all places from database
func GetAllCategoriesHandler(c *gin.Context) {
	var categories []models.Category
	categories = orm.GetCategories()
	if categories != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": categories})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": []models.Category{}})
	}
}

// InsertCategoryHandler insert place into database
func InsertCategoryHandler(c *gin.Context) {
	title := c.PostForm("title")
	if title != "" {
		err := orm.InsertCategory(models.Category{Title: title})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		}
		c.JSON(http.StatusCreated, gin.H{"code": http.StatusCreated, "message": "Category has been inserted to database."})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Title cannot be empty."})
	}
}
