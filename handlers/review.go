package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rakawestu/explore-jogja-api/models"
	"github.com/rakawestu/explore-jogja-api/orm"
)

// GetPlaceReviews returns all reviews for specific place
func GetPlaceReviews(c *gin.Context) {
	placeID := c.Param("id")
	limitQuery := c.Query("limit")
	offsetQuery := c.Query("offset")

	var limit int64 = 10
	var offset int64

	if limitQuery != "" {
		limit1, _ := strconv.ParseInt(limitQuery, 10, 32)
		limit = limit1
	}
	if offsetQuery != "" {
		offset1, _ := strconv.ParseInt(offsetQuery, 0, 32)
		offset = offset1
	}
	reviews, err := orm.GetPlaceReviews(placeID, int(limit), int(offset))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusOK, "message": err.Error()})
	} else {
		if reviews == nil {
			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": []models.Review{}})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": reviews})
		}
	}
}

// InsertReviewForPlace inserts a review into specific place
func InsertReviewForPlace(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	placeID := c.Param("id")
	user := c.PostForm("user")

	var username string
	if user != "" {
		username = user
	} else {
		username = "anonym"
	}

	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Title cannot be empty."})
	} else if content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Content cannot be empty."})
	} else {
		err := orm.InsertPlaceReview(models.Review{Title: title, Content: content, PlaceID: placeID, User: username})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		} else {
			c.JSON(http.StatusCreated, gin.H{"code": http.StatusCreated, "message": "Review has been inserted to database."})
		}
	}
}
