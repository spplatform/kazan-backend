package handlers

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/go-openapi/strfmt"
	"github.com/spplatform/kazan-backend/models"
	"github.com/spplatform/kazan-backend/persistence/entity"
	"github.com/spplatform/kazan-backend/restapi/operations/coupon"
	"github.com/spplatform/kazan-backend/restapi/operations/order"
	"github.com/spplatform/kazan-backend/restapi/operations/route"
	"log"
	"strings"
	"time"

	"github.com/globalsign/mgo"
	"github.com/go-openapi/runtime/middleware"
)

type Handler struct {
	db *mgo.Database
}

func NewHandler(user, pwd, host, database string) (*Handler, error) {
	conURI := fmt.Sprintf("mongodb://%s:%s@%s/%s", user, pwd, host, database)
	time.Sleep(5 * time.Second)
	log.Printf("connect to MongoDB at [%s]", conURI)
	session, err := mgo.Dial(conURI)
	if err != nil {
		return nil, err
	}

	return &Handler{
		db: session.DB(database),
	}, nil
}

func (h *Handler) HandleGetOrder(p order.GetOrderIDParams) middleware.Responder {
	log.Printf("HandleGetOrder [%s]", p.ID)

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
	return order.NewDeleteOrderIDAccepted().WithPayload(&models.StatusResponse{
		Message: "Canceled",
	})
}

func (h *Handler) HandlePostOrder(p order.PostOrderParams) middleware.Responder {
	log.Print("HandlePostOrder")

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

	id := o.ID.Hex()
	payURL := generatePaymentUrl(id)
	resp := models.OrderCreateResponse{
		ID:         &id,
		PaymentURL: &payURL,
		Status:     &o.Status,
		Coupon:     o.Coupon,
		Positions:  make([]*models.OrderItem, 0, len(o.Items)),
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

			for _, cpos := range cafe.Positions {
				rCafe.Positions = append(rCafe.Positions, &models.CafeDishResponse{
					ID:       &cpos.ID,
					ImageURL: &cpos.ImageURL,
					Name:     &cpos.Name,
					Price:    &cpos.Price,
				})
			}

			rStop.Cafes = append(rStop.Cafes, &rCafe)
		}
		resp.Stops = append(resp.Stops, &rStop)
	}

	return route.NewGetTicketIDRouteOK().WithPayload(&resp)
}

func (h *Handler) HandleGetCoupon(p coupon.GetCouponIDParams) middleware.Responder {
	log.Printf("HandleGetCoupon [%s]", p.ID)

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

func generatePaymentUrl(id string) string {
	return "/pay/" + id
}
