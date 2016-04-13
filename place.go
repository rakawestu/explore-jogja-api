package main

import "gopkg.in/mgo.v2/bson"

// Place object
type Place struct {
	Title       string        `form:"title" json:"title" binding:"required"`
	Location    Location      `form:"location" json:"location" binding:"required"`
	Description string        `form:"description" json:"description" binding:"required"`
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
}
