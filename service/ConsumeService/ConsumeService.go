package ConsumeService

import (
	"github.com/rabbitmq/amqp091-go"
	"learn_rabbitmq/model/appModel"
	"log"
	"net/http"
	"strings"
)

func ConsumeService(request *http.Request) (status int, err error) {
	var (
		channel = appModel.RabbitChannel
	)
	rawQuery := strings.Split(request.URL.RawQuery, "=")
	if len(rawQuery) < 1 {
		return
	}
	declare, err := channel.QueueDeclare(rawQuery[1], false, false, false, false, nil)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	message, err := channel.Consume(declare.Name, "", true, false, false, false, nil)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	go listenFromChannel(message)

	return http.StatusOK, nil
}

func listenFromChannel(msg <-chan amqp091.Delivery) {
	for data := range msg {
		log.Printf(`Message : %s`, data.Body)
	}
}
