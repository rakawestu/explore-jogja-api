package main

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	collectionName = "places"
)

// GetPlaces is a function to get all place from database
func GetPlaces(limit int, skip int) []Place {
	session, err := mgo.Dial(mongoDBUrl)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	var places []Place
	c := session.DB(mongoDBName).C(collectionName)

	err1 := c.Find(bson.M{}).Limit(limit).Skip(skip).All(&places)
	if err1 != nil {
		panic(err1)
	}
	return places
}

// GetPlacesBasedOnCategory get places data based on category
func GetPlacesBasedOnCategory(category string, limit int, skip int) []Place {
	session, err := mgo.Dial(mongoDBUrl)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	var places []Place
	c := session.DB(mongoDBName).C(collectionName)

	err1 := c.Find(bson.M{"category": category}).Limit(limit).Skip(skip).All(&places)
	if err1 != nil {
		panic(err1)
	}
	return places
}

// InsertPlace is a function to insert place into database
func InsertPlace(place Place) error {
	session, err := mgo.Dial(mongoDBUrl)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB(mongoDBName).C(collectionName)

	err1 := c.Insert(&place)
	return err1
}

// GetSinglePlace is a function to get place based on ID
func GetSinglePlace(id string) (Place, error) {
	session, err := mgo.Dial(mongoDBUrl)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	var place Place

	if !bson.IsObjectIdHex(id) {
		return place, errors.New("Place ID is not valid.")
	}

	c := session.DB(mongoDBName).C(collectionName)

	err1 := c.FindId(bson.ObjectIdHex(id)).One(&place)
	return place, err1
}

// DeletePlace is a function to delete specific place based on ID
func DeletePlace(id string) error {
	session, err := mgo.Dial(mongoDBUrl)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	if !bson.IsObjectIdHex(id) {
		return errors.New("Place ID is not valid.")
	}

	c := session.DB(mongoDBName).C(collectionName)

	err1 := c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err1
}

// UpdatePlace is a function to update place data
func UpdatePlace(id string, place Place) error {
	session, err := mgo.Dial(mongoDBUrl)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	currentPlace, err1 := GetSinglePlace(id)
	if err1 != nil {
		return err1
	}

	if place.Title != "" {
		currentPlace.Title = place.Title
	}

	if place.Location.Address != "" {
		currentPlace.Location.Address = place.Location.Address
	}

	if place.Description != "" {
		currentPlace.Description = place.Description
	}

	if place.Category != "" {
		currentPlace.Category = place.Category
	}

	if place.Location.Latitude != 0 {
		currentPlace.Location.Latitude = place.Location.Latitude
	}

	if place.Location.Longitude != 0 {
		currentPlace.Location.Longitude = place.Location.Longitude
	}

	if place.OpeningHours != "" {
		currentPlace.OpeningHours = place.OpeningHours
	}

	if place.PriceRange != "" {
		currentPlace.PriceRange = place.PriceRange
	}

	c := session.DB(mongoDBName).C(collectionName)

	_, err2 := c.UpsertId(bson.ObjectIdHex(id), currentPlace)

	return err2
}
