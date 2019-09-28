package entity

import (
	"github.com/globalsign/mgo/bson"
)

type SenderAccount struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	Active     bool          `bson:"active"`
	ServerHost string        `bson:"server_host"`
	ServerPort string        `bson:"server_port"`
	Email      string        `bson:"email"`
	Password   string        `bson:"password"`
}

type Reporter struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	EventType string        `bson:"event_type"`
	Email     string        `bson:"email"`
}
