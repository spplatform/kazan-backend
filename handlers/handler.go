package handlers

import (
	"github.com/globalsign/mgo/bson"
	"github.com/go-openapi/strfmt"
	"github.com/spplatform/kazan-backend/models"
	"github.com/spplatform/kazan-backend/persistence/entity"
	"github.com/spplatform/kazan-backend/reporting"
	"github.com/spplatform/kazan-backend/restapi/operations/coupon"
	"github.com/spplatform/kazan-backend/restapi/operations/order"
	"github.com/spplatform/kazan-backend/restapi/operations/payment"
	"github.com/spplatform/kazan-backend/restapi/operations/route"
	"log"
	"strings"
	"time"

	"github.com/globalsign/mgo"
	"github.com/go-openapi/runtime/middleware"
)

type Reporter interface {
	Push(e reporting.Event)
}

type Handler struct {
	db *mgo.Database
	r  Reporter
}

func NewHandler(db *mgo.Database, r Reporter) *Handler {
	return &Handler{
		db: db,
		r:  r,
	}
}

func (h *Handler) HandleGetOrder(p order.GetOrderIDParams) middleware.Responder {
	log.Printf("HandleGetOrder [%s]", p.ID)
	defer func(t time.Time) {
		log.Printf("HandleGetOrder took %fs", time.Since(t).Seconds())
	}(time.Now())

	result := entity.Order{}
	err := h.db.C(entity.CollectionOrder).FindId(bson.ObjectIdHex(p.ID)).One(&result)
	if err == mgo.ErrNotFound {
		return order.NewGetOrderIDNotFound().WithPayload(&models.StatusResponse{
			Message: err.Error(),
		})
	} else if err != nil {
		return order.NewGetOrderIDInternalServerError().WithPayload(&models.StatusResponse{
			Message: err.Error(),
		})
	}

	id := result.ID.Hex()
	resp := models.OrderResponse{
		ID:        &id,
		Status:    &result.Status,
		Positions: make([]*models.OrderItem, 0, len(result.Items)),
	}

	for _, item := range result.Items {
		resp.Positions = append(resp.Positions, &models.OrderItem{
			ID:         &item.PositionID,
			Amount:     &item.Amount,
			TotalPrice: item.Total,
		})
	}

	return order.NewGetOrderIDOK().WithPayload(&resp)
}

func (h *Handler) HandleDeleteOrder(p order.DeleteOrderIDParams) middleware.Responder {
	log.Printf("HandleDeleteOrder [%s]", p.ID)
	defer func(t time.Time) {
		log.Printf("HandleDeleteOrder took %fs", time.Since(t).Seconds())
	}(time.Now())

	err := h.db.C(entity.CollectionOrder).UpdateId(
		bson.ObjectIdHex(p.ID),
		bson.M{"$set": bson.M{"status": entity.OrderStatusCanceled}})
	if err == mgo.ErrNotFound {
		return order.NewDeleteOrderIDNotFound().WithPayload(&models.StatusResponse{
			Message: err.Error(),
		})
	} else if err != nil {
		return order.NewDeleteOrderIDInternalServerError().WithPayload(&models.StatusResponse{
			Message: err.Error(),
		})
	}

	err = h.db.C(entity.CollectionPayment).UpdateId(
		bson.M{"order_id": bson.ObjectIdHex(p.ID)},
		bson.M{"$set": bson.M{"status": entity.PaymentStatusCanceled}})
	if err != nil {
		return order.NewDeleteOrderIDInternalServerError().WithPayload(&models.StatusResponse{
			Message: err.Error(),
		})
	}

	return order.NewDeleteOrderIDAccepted().WithPayload(&models.StatusResponse{
		Message: "Canceled",
	})
}

func (h *Handler) HandlePostOrder(p order.PostOrderParams) middleware.Responder {
	log.Print("HandlePostOrder")
	defer func(t time.Time) {
		log.Printf("HandlePostOrder took %fs", time.Since(t).Seconds())
	}(time.Now())

	o := entity.Order{
		ID:     bson.NewObjectId(),
		UserID: *p.Body.UserID,
		CafeID: bson.ObjectIdHex(*p.Body.Order.CafeID),
		Status: entity.OrderStatusNew,
		Coupon: p.Body.Order.Coupon,
		Items:  make([]entity.OrderItem, 0, len(p.Body.Order.Positions)),
	}

	//get cafe prices
	cafe := entity.Cafe{}
	err := h.db.C(entity.CollectionCafe).FindId(o.CafeID).One(&cafe)
	if err != nil {
		return order.NewPostOrderBadRequest().WithPayload(&models.StatusResponse{
			Message: err.Error(),
		})
	}

	prices := make(map[string]int64, len(cafe.Positions))
	for _, pos := range cafe.Positions {
		prices[pos.ID] = pos.Price
	}

	//calculate coupon discount
	discount := 1.0
	if p.Body.Order.Coupon != "" {
		c := entity.Coupon{}
		err := h.db.C(entity.CollectionCoupon).Find(bson.M{"code": strings.ToLower(p.Body.Order.Coupon)}).One(&c)
		if err == nil && c.IsValid(time.Now()) {
			log.Printf("redeem coupon %s", p.Body.Order.Coupon)
			discount -= float64(c.Discount)
			if c.Remaining != nil {
				_ = h.db.C(entity.CollectionCoupon).Update(
					bson.M{"code": strings.ToLower(p.Body.Order.Coupon)},
					bson.M{"$inc": bson.M{"remaining": -1}})
			}
		} else {
			log.Printf("coupon %s is not valid", p.Body.Order.Coupon)
		}
	}

	for _, pos := range p.Body.Order.Positions {
		o.Items = append(o.Items, entity.OrderItem{
			PositionID: *pos.ID,
			Amount:     *pos.Amount,
			Total:      int64(float64(*pos.Amount*prices[*pos.ID]) * discount),
		})
	}

	err = h.db.C(entity.CollectionOrder).Insert(&o)
	if err != nil {
		return order.NewPostOrderInternalServerError().WithPayload(&models.StatusResponse{
			Message: err.Error(),
		})
	}

	pay := entity.Payment{
		ID:      bson.NewObjectId(),
		OrderID: o.ID,
		UserID:  o.UserID,
		Status:  entity.PaymentStatusNew,
	}
	err = h.db.C(entity.CollectionPayment).Insert(&pay)
	if err != nil {
		return order.NewPostOrderInternalServerError().WithPayload(&models.StatusResponse{
			Message: err.Error(),
		})
	}

	id := o.ID.Hex()
	pid := pay.ID.Hex()
	resp := models.OrderCreateResponse{
		ID:        &id,
		PaymentID: &pid,
		Status:    &o.Status,
		Coupon:    o.Coupon,
		Positions: make([]*models.OrderItem, 0, len(o.Items)),
	}

	for _, item := range o.Items {
		resp.Positions = append(resp.Positions, &models.OrderItem{
			ID:         &item.PositionID,
			Amount:     &item.Amount,
			TotalPrice: item.Total,
		})
	}

	return order.NewPostOrderCreated().WithPayload(&resp)
}

func (h *Handler) HandleGetTicketRoute(p route.GetTicketIDRouteParams) middleware.Responder {
	log.Printf("HandleGetTicketRoute [%s]", p.ID)
	defer func(t time.Time) {
		log.Printf("HandleGetTicketRoute took %fs", time.Since(t).Seconds())
	}(time.Now())

	result := entity.Route{}
	err := h.db.C(entity.CollectionRoute).Find(bson.M{"tickets": p.ID}).One(&result)
	if err == mgo.ErrNotFound {
		return route.NewGetTicketIDRouteNotFound().WithPayload(&models.StatusResponse{
			Message: err.Error(),
		})
	} else if err != nil {
		return route.NewGetTicketIDRouteInternalServerError().WithPayload(&models.StatusResponse{
			Message: err.Error(),
		})
	}

	resp := models.RouteResponse{
		TrainNumber: &result.TrainNumber,
		Stops:       make([]*models.RouteResponseStopsItems0, 0, len(result.Stops)),
	}

	for _, stop := range result.Stops {
		city := entity.City{}
		name := "-"
		err := h.db.C(entity.CollectionCity).FindId(stop.CityID).One(&city)
		if err == nil {
			name = city.Name
		}

		cid := stop.CityID.Hex()
		dt := strfmt.DateTime(stop.DateTime.Time())
		rStop := models.RouteResponseStopsItems0{
			CityID:   &cid,
			DateTime: &dt,
			Duration: &stop.Duration,
			Name:     &name,
		}

		cafes := []*entity.Cafe{}
		_ = h.db.C(entity.CollectionCafe).Find(bson.M{"city_id": stop.CityID}).All(&cafes)
		rStop.Cafes = make([]*models.CafeResponse, 0, len(cafes))

		for _, cafe := range cafes {

			cfid := cafe.ID.Hex()
			rCafe := models.CafeResponse{
				CityID:    &cid,
				ID:        &cfid,
				Name:      &cafe.Name,
				Positions: make([]*models.CafeDishResponse, 0, len(cafe.Positions)),
			}

			for i := range cafe.Positions {
				rPos := models.CafeDishResponse{
					ID:          &cafe.Positions[i].ID,
					Price:       &cafe.Positions[i].Price,
					ImageURL:    &cafe.Positions[i].ImageURL,
					Name:        &cafe.Positions[i].Name,
					Description: cafe.Positions[i].Description,
				}
				rCafe.Positions = append(rCafe.Positions, &rPos)
			}

			rStop.Cafes = append(rStop.Cafes, &rCafe)
		}
		resp.Stops = append(resp.Stops, &rStop)
	}

	return route.NewGetTicketIDRouteOK().WithPayload(&resp)
}

func (h *Handler) HandleGetCoupon(p coupon.GetCouponIDParams) middleware.Responder {
	log.Printf("HandleGetCoupon [%s]", p.ID)
	defer func(t time.Time) {
		log.Printf("HandleGetCoupon took %fs", time.Since(t).Seconds())
	}(time.Now())

	c := entity.Coupon{}
	err := h.db.C(entity.CollectionCoupon).Find(bson.M{"code": strings.ToLower(p.ID)}).One(&c)
	if err == mgo.ErrNotFound {
		return coupon.NewGetCouponIDNotFound().WithPayload(&models.StatusResponse{
			Message: err.Error(),
		})
	} else if err != nil {
		return coupon.NewGetCouponIDInternalServerError().WithPayload(&models.StatusResponse{
			Message: err.Error(),
		})
	}

	valid := c.IsValid(time.Now())

	resp := models.CouponResponse{
		Discount: c.Discount,
		Valid:    &valid,
	}

	return coupon.NewGetCouponIDOK().WithPayload(&resp)
}

func (h *Handler) HandlePutPay(p payment.PutPayParams) middleware.Responder {
	log.Printf("HandlePutPay [%s]", *p.Body.PaymentID)
	defer func(t time.Time) {
		log.Printf("HandlePutPay took %fs", time.Since(t).Seconds())
	}(time.Now())

	pay := entity.Payment{}
	err := h.db.C(entity.CollectionPayment).FindId(bson.ObjectIdHex(*p.Body.PaymentID)).One(&pay)
	if err == mgo.ErrNotFound {
		return payment.NewPutPayNotFound().WithPayload(&models.StatusResponse{
			Message: err.Error(),
		})
	} else if err != nil {
		return payment.NewPutPayInternalServerError().WithPayload(&models.StatusResponse{
			Message: err.Error(),
		})
	}

	//validate
	if pay.UserID != *p.Body.UserID {
		return payment.NewPutPayBadRequest().WithPayload(&models.StatusResponse{
			Message: "wrong user id",
		})
	}

	if pay.Status != entity.PaymentStatusNew {
		return payment.NewPutPayBadRequest().WithPayload(&models.StatusResponse{
			Message: "payment with status '" + pay.Status + "' cannot be processed",
		})
	}

	pay.Status = entity.PaymentStatusPaid
	err = h.db.C(entity.CollectionPayment).Update(bson.M{"_id": pay.ID}, bson.M{"$set": pay})
	if err != nil {
		return payment.NewPutPayInternalServerError().WithPayload(&models.StatusResponse{
			Message: err.Error(),
		})
	}

	err = h.db.C(entity.CollectionOrder).Update(
		bson.M{"_id": pay.OrderID},
		bson.M{"$set": bson.M{"status": entity.OrderStatusPaid}})
	if err != nil {
		return payment.NewPutPayInternalServerError().WithPayload(&models.StatusResponse{
			Message: err.Error(),
		})
	}

	go h.r.Push(reporting.Event{
		Type:    reporting.EventTypePayment,
		Payload: pay,
	})

	resp := models.PaymentResponse{
		Status: &pay.Status,
	}

	return payment.NewPutPayOK().WithPayload(&resp)
}
