package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rakawestu/explore-jogja-api/handlers"
)

// CategoryRoutes is router for category function
func CategoryRoutes(router *gin.RouterGroup) {
	router.GET("/categories", handlers.GetAllCategoriesHandler)
	router.POST("/categories", handlers.InsertCategoryHandler)
}
