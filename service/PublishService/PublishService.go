package PublishService

import (
	"encoding/json"
	"github.com/rabbitmq/amqp091-go"
	"learn_rabbitmq/model/appModel"
	"net/http"
)

type PublishRequest struct {
	QueueName string `json:"queue_name"`
	Message   string `json:"message"`
}

func PublishService(request *http.Request) (status int, err error) {
	var (
		requestStruct PublishRequest
		channel       = appModel.RabbitChannel
	)
	err = json.NewDecoder(request.Body).Decode(&requestStruct)
	defer request.Body.Close()
	if err != nil {
		status = http.StatusInternalServerError
		return
	}

	declare, err := channel.QueueDeclare(requestStruct.QueueName, false, false, false, false, nil)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	body, _ := json.Marshal(&requestStruct)

	err = channel.Publish("", declare.Name, false, false, amqp091.Publishing{
		ContentType: "application/json",
		Body:        body,
	})

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
