package entity

import "github.com/globalsign/mgo/bson"

type Cafe struct {
	ID           bson.ObjectId  `bson:"_id,omitempty"`
	Name         string         `bson:"name"`
	Rating       float32        `bson:"rating"`
	Cuisine      string         `bson:"cuisine"`
	MinimumPrice int64          `bson:"minimum_price"`
	CityID       bson.ObjectId  `bson:"city_id"`
	Positions    []CafePosition `bson:"positions"`
}

type CafePosition struct {
	ID          string `bson:"id"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
	Price       int64  `bson:"price"`
	ImageURL    string `bson:"image_url"`
}
