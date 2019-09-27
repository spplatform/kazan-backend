package entity

import "github.com/globalsign/mgo/bson"

type City struct {
	CityID bson.ObjectId `bson:"_id,omitempty"`
	Name   string        `bson:"name"`
}
