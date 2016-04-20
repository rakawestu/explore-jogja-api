package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rakawestu/explore-jogja-api/models"
	"github.com/rakawestu/explore-jogja-api/orm"
)

// GetAllPlacesHandler returns all places from database
func GetAllPlacesHandler(c *gin.Context) {
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

	var places []models.Place
	category := c.Query("category")
	if category != "" {
		places = orm.GetPlacesBasedOnCategory(strings.ToLower(category), int(limit), int(offset))
	} else {
		places = orm.GetPlaces(int(limit), int(offset))
	}
	if places != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": places})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": []models.Place{}})
	}
}

// GetSinglePlaceHandler returns place using specific identifier
func GetSinglePlaceHandler(c *gin.Context) {
	id := c.Param("id")
	place, err := orm.GetSinglePlace(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": place})
	}
}

// InsertPlaceHandler insert place into database
func InsertPlaceHandler(c *gin.Context) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	address := c.PostForm("address")
	lat := c.PostForm("latitude")
	lng := c.PostForm("longitude")
	category := c.PostForm("category")
	opHours := c.PostForm("opening_hours")
	priceRange := c.PostForm("price_range")

	latitude, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Latitude is not a valid number."})
	}
	longitude, err1 := strconv.ParseFloat(lng, 64)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Longitude is not a valid number."})
	}

	err2 := orm.InsertPlace(models.Place{Title: title, Description: description, Location: models.Location{Latitude: latitude, Longitude: longitude, Address: address}, Category: category, OpeningHours: opHours, PriceRange: priceRange})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err2.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"code": http.StatusCreated, "message": "Place has been inserted to database."})
}

// DeletePlaceHandler delete place from database
func DeletePlaceHandler(c *gin.Context) {
	id := c.Param("id")
	err := orm.DeletePlace(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Place has been deleted."})
	}
}

// UpdatePlaceHandler updates selected place
func UpdatePlaceHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.PostForm("title")
	address := c.PostForm("address")
	description := c.PostForm("description")
	lat := c.PostForm("latitude")
	lng := c.PostForm("longitude")
	category := c.PostForm("category")
	opHours := c.PostForm("opening_hours")
	priceRange := c.PostForm("price_range")

	var latitude float64
	var longitude float64
	if lat != "" {
		latitude, _ = strconv.ParseFloat(lat, 64)
	} else {
		latitude = 0
	}

	if lng != "" {
		longitude, _ = strconv.ParseFloat(lng, 64)
	} else {
		longitude = 0
	}

	err2 := orm.UpdatePlace(id, models.Place{Title: title, Description: description, Location: models.Location{Latitude: latitude, Longitude: longitude, Address: address}, Category: category, OpeningHours: opHours, PriceRange: priceRange})
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err2.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Place has been updated."})
	}
}
