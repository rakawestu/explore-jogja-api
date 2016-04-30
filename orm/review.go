package orm

import (
	"errors"

	"github.com/rakawestu/explore-jogja-api/models"
	"gopkg.in/mgo.v2/bson"
)

const (
	collectionNameReview = "reviews"
)

// GetPlaceReviews is a function to get reviews on specific place
func GetPlaceReviews(placeID string, limit int, offset int) ([]models.Review, error) {
	var reviews []models.Review
	if !bson.IsObjectIdHex(placeID) {
		return reviews, errors.New("Place ID is not valid.")
	}
	c := MongoDB.C(collectionNameReview)
	err := c.Find(bson.M{"placeid": placeID}).Limit(limit).Skip(offset).All(&reviews)

	return reviews, err
}

// InsertPlaceReview is a function to insert review on specific place
func InsertPlaceReview(review models.Review) error {
	c := MongoDB.C(collectionNameReview)

	err1 := c.Insert(&review)
	return err1
}
