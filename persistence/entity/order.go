package entity

import "github.com/globalsign/mgo/bson"

type Order struct {
	ID     bson.ObjectId `bson:"_id,omitempty"`
	UserID string        `bson:"user_id"`
	CafeID bson.ObjectId `bson:"cafe_id"`
	Status string        `bson:"status"`
	Coupon string        `bson:"coupon,omitempty"`
	Items  []OrderItem   `bson:"items"`
}
type OrderItem struct {
	PositionID string `bson:"position_id"`
	Amount     int64  `bson:"amount"`
	Total      int64  `bson:"total_price"`
}

const (
	OrderStatusNew      = "new"
	OrderStatusPaid     = "paid"
	OrderStatusDelivery = "delivery"
	OrderStatusDone     = "done"
	OrderStatusCanceled = "canceled"
)
