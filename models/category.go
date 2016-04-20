package models

import "gopkg.in/mgo.v2/bson"

// Category is model class for category
type Category struct {
	ID    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title string        `form:"title" json:"title" binding:"required"`
}
