package models

import "gopkg.in/mgo.v2/bson"

// Place object
type Place struct {
	Title        string        `form:"title" json:"title" binding:"required"`
	Location     Location      `form:"location" json:"location" binding:"required"`
	Description  string        `form:"description" json:"description" binding:"required"`
	Category     string        `form:"category" json:"category"`
	OpeningHours string        `form:"opening_hours" json:"opening_hours"`
	PriceRange   string        `form:"price_range" json:"price_range"`
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
}
