package handlers

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/spplatform/kazan-backend/models"
	"github.com/spplatform/kazan-backend/persistence/entity"
	"github.com/spplatform/kazan-backend/restapi/operations/order"
	"github.com/spplatform/kazan-backend/restapi/operations/route"
	"log"
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
		return order.NewGetOrderIDNotFound().WithPayload(&models.ErrorResponse{
			Message: err.Error(),
		})
	} else if err != nil {
		return order.NewGetOrderIDInternalServerError().WithPayload(&models.ErrorResponse{
			Message: err.Error(),
		})
	}

	id := result.ID.Hex()
	resp := models.OrderResponse{
		ID: &id,
		Status: &models.OrderStatusResponse{
			Description: &result.Status.Status,
			Status:      &result.Status.Description,
		},
		Positions: make([]*models.OrderItem, 0, len(result.Items)),
	}

	for _, item := range result.Items {
		resp.Positions = append(resp.Positions, &models.OrderItem{
			ID:     &item.PositionID,
			Amount: &item.Amount,
		})
	}

	return order.NewGetOrderIDOK().WithPayload(&resp)
}

func (h *Handler) HandlePostOrder(order.PostOrderParams) middleware.Responder {
	log.Print("HandlePostOrder")
	return middleware.NotImplemented("operation HandlePostOrder is ok!")
}

func (h *Handler) HandleGetTicketRoute(route.GetTicketIDRouteParams) middleware.Responder {
	log.Print("HandleGetTicketRoute")
	return middleware.NotImplemented("operation HandleGetTicketRoute is ok!")
}
