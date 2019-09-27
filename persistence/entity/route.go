package entity

import (
	"github.com/globalsign/mgo/bson"
)

type Route struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	TrainNumber string        `bson:"train_number"`
	Stops       []RouteStop   `bson:"stops"`
	Tickets     []string      `bson:"tickets"`
}

type RouteStop struct {
	CityID   bson.ObjectId       `bson:"city_id"`
	DateTime bson.MongoTimestamp `bson:"date_time"`
	Duration int64               `bson:"duration"`
}
