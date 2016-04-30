package models

import "gopkg.in/mgo.v2/bson"

// Review object
type Review struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title   string        `form:"title" json:"title" bson:"title" binding:"required"`
	Content string        `form:"content" json:"content" bson:"content" binding:"required"`
	User    string        `form:"user" json:"user" bson:"user" binding:"required"`
	PlaceID string        `form:"place_id" json:"place_id" bson:place_id`
}
