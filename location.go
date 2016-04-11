package main

// Location is model of location
type Location struct {
	Address   string  `form:"address" json:"address" binding:"required"`
	Latitude  float64 `form:"latitude" json:"latitude" binding:"required"`
	Longitude float64 `form:"longitude" json:"longitude" binding:"required"`
}
