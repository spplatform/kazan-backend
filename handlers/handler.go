package handlers

import (
	"fmt"
	"log"
	"time"

	"github.com/globalsign/mgo"
	"github.com/go-openapi/runtime/middleware"
	"github.com/spplatform/kazan-backend/restapi/operations/test"
)

type Handler struct {
	mgoSession *mgo.Session
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
		mgoSession: session,
	}, nil
}

func (h *Handler) HandleTest(params test.GetTestParams) middleware.Responder {
	log.Print("HandleTest")
	return middleware.NotImplemented("operation test.GetTest is ok!")
}
