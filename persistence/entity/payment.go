package entity

import "github.com/globalsign/mgo/bson"

type Payment struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	OrderID bson.ObjectId `bson:"order_id"`
	UserID  string        `bson:"user_id"`
	Status  string        `bson:"status"`
}

const (
	PaymentStatusNew      = "new"
	PaymentStatusPaid     = "paid"
	PaymentStatusCanceled = "canceled"
)
