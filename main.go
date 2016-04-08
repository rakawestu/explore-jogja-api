package main

import (
	"flag"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	// MongoDBUrlKey is flag key for Mongo DB URL
	MongoDBUrlKey = "MongoDBUrl"
	// MongoDBNameKey is flag key for Mongo DB name
	MongoDBNameKey = "MongoDBName"
	// MongoDBUrlDefaultValue is Mongo DB default URL
	MongoDBUrlDefaultValue = "localhost"
	// MongoDBNameDefaultValue is Mongo DB default name
	MongoDBNameDefaultValue = "explore-jogja"
)

var mongoDBUrl, mongoDBName string

func main() {
	router := gin.Default()

	flag.StringVar(&mongoDBUrl, MongoDBUrlKey, MongoDBUrlDefaultValue, "Mongo DB URL.")
	flag.StringVar(&mongoDBName, MongoDBNameKey, MongoDBNameDefaultValue, "Mongo DB database name")

	router.GET("/places", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": GetPlaces()})
	})

	router.GET("/places/:id", func(c *gin.Context) {
		id := c.Param("id")
		place, err := GetSinglePlace(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": place})
		}
	})

	router.POST("/places", func(c *gin.Context) {
		title := c.PostForm("title")
		description := c.PostForm("description")
		address := c.PostForm("address")
		lat := c.PostForm("latitude")
		lng := c.PostForm("longitude")

		latitude, err := strconv.ParseFloat(lat, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Latitude is not a valid number."})
		}
		longitude, err1 := strconv.ParseFloat(lng, 64)
		if err1 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Longitude is not a valid number."})
		}

		err2 := InsertPlace(Place{
			Title: title, Latitude: latitude, Longitude: longitude, Description: description, Address: address})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err2.Error()})
		}

		c.JSON(http.StatusCreated, gin.H{"code": http.StatusCreated, "message": "Place has been inserted to database."})

	})

	router.DELETE("/places/:id", func(c *gin.Context) {
		id := c.Param("id")
		err := DeletePlace(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Place has been deleted."})
		}
	})

	router.PUT("/places/:id", func(c *gin.Context) {
		id := c.Param("id")
		title := c.PostForm("title")
		address := c.PostForm("address")
		description := c.PostForm("description")
		lat := c.PostForm("latitude")
		lng := c.PostForm("longitude")

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

		err2 := UpdatePlace(id, Place{Title: title, Latitude: latitude, Longitude: longitude, Description: description, Address: address})
		if err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err2.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Place has been updated."})
		}
	})

	router.Run(":8080")
}
