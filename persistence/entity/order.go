package entity

import "github.com/globalsign/mgo/bson"

type Order struct {
	ID     bson.ObjectId `bson:"_id"`
	UserID string        `bson:"user_id"`
	CafeID bson.ObjectId `bson:"cafe_id"`
	Status OrderStatus   `bson:"status"`
	Items  []OrderItem   `bson:"items"`
}
type OrderItem struct {
	PositionID string `bson:"position_id"`
	Amount     int64  `bson:"amount"`
}

type OrderStatus struct {
	Status      string `bson:"status"`
	Description string `bson:"description"`
}
