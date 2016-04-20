package orm

import (
	"github.com/rakawestu/explore-jogja-api/models"
	"gopkg.in/mgo.v2/bson"
)

const (
	collectionNameCategory = "categories"
)

// GetCategories is orm method for get all categories from database
func GetCategories() []models.Category {

	var categories []models.Category
	c := MongoDB.C(collectionNameCategory)

	err1 := c.Find(bson.M{}).All(&categories)
	if err1 != nil {
		panic(err1)
	}
	return categories
}

// InsertCategory is orm method for inserting category into database
func InsertCategory(category models.Category) error {

	c := MongoDB.C(collectionNameCategory)

	err1 := c.Insert(&category)
	return err1
}
