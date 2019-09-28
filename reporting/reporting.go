package reporting

import (
	"context"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/spplatform/kazan-backend/persistence/entity"
	"log"
	"net/smtp"
	"time"
)

const (
	EventTypePayment = "payment"
)

type Event struct {
	Type    string
	Payload interface{}
}

func RunReportingJob(ctx context.Context, db *mgo.Database) *reporter {
	r := reporter{
		db:        db,
		eventChan: make(chan Event),
	}
	go r.handle(ctx)

	return &r

}

type reporter struct {
	db        *mgo.Database
	eventChan chan Event
}

func (r *reporter) Push(e Event) {
	r.eventChan <- e
}

func (r *reporter) handle(ctx context.Context) {
	log.Print("start report job")
	defer log.Print("finish report job")

	for {
		select {
		case <-ctx.Done():
			return
		case e, ok := <-r.eventChan:
			if !ok {
				return
			}
			r.handleEvent(e)
		}
	}
}

func (r *reporter) handleEvent(e Event) {
	switch e.Type {
	case EventTypePayment:
		r.handlePayment(e.Payload.(entity.Payment))
	}
}

func (r *reporter) handlePayment(p entity.Payment) {
	log.Printf("handlePayment [%s]", p.ID.Hex())
	defer func(t time.Time) {
		log.Printf("handlePayment took %fs", time.Since(t).Seconds())
	}(time.Now())

	sender := entity.SenderAccount{}
	err := r.db.C(entity.CollectionSender).Find(bson.M{"active": true}).One(&sender)
	if err != nil {
		log.Printf("error during reporting: %v", err)
		return
	}

	receivers := []*entity.Reporter{}
	err = r.db.C(entity.CollectionReporter).Find(bson.M{"event_type": EventTypePayment}).All(&receivers)
	if err != nil {
		log.Printf("error during reporting: %v", err)
		return
	}

	rcvList := make([]string, 0, len(receivers))

	for _, r := range receivers {
		rcvList = append(rcvList, r.Email)
	}

	if len(rcvList) == 0 {
		return
	}
	msg := fmt.Sprintf("Subject:Отчет о платеже FOR\n\nОтчет\n\nПлатеж %s выполнен в %s.\n\nКоманда Food on Rails", p.ID.Hex(), time.Now().Format("02.01.06 15:04"))

	err = smtp.SendMail(sender.ServerHost+":"+sender.ServerPort,
		smtp.PlainAuth("", sender.Email, sender.Password, sender.ServerHost),
		sender.Email, rcvList, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
	}
}
