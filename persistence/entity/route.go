package entity

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type Route struct {
}

type RouteStop struct {
	CityID   bson.ObjectId `bson:"city_id"`
	DateTime time.Time     `bson:"date_time"`
	Duration int64         `bson:"duration"`
}
