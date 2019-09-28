package mongo

import (
	"fmt"
	"github.com/globalsign/mgo"
	"log"
	"time"
)

func Connect(user, pwd, host, database string) (*mgo.Database, error) {
	conURI := fmt.Sprintf("mongodb://%s:%s@%s/%s", user, pwd, host, database)
	time.Sleep(5 * time.Second)
	log.Printf("connect to MongoDB at [%s]", conURI)
	session, err := mgo.Dial(conURI)
	if err != nil {
		return nil, err
	}
	return session.DB(database), nil
}
