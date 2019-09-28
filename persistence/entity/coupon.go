package entity

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type Coupon struct {
	Code       string              `bson:"code"`
	Discount   float32             `bson:"discount"`
	ValidUntil bson.MongoTimestamp `bson:"valid_until"`
	Remaining  *int64              `bson:"remaining,omitempty"`
}

func (c Coupon) IsValid(t time.Time) bool {
	return c.ValidUntil.Time().After(t) && (c.Remaining != nil && *c.Remaining > 0) || c.Remaining == nil
}
