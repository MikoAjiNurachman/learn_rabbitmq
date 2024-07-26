package appModel

import (
	"github.com/rabbitmq/amqp091-go"
	"learn_rabbitmq/model/errorModel"
	"log"
	"os"
	"time"
)

var RabbitChannel *amqp091.Channel

func GenerateAppMxodel() {
	rabbitUrl := os.Getenv("RABBITMQ_URL")

	conn, err := connectRabbitMQ(rabbitUrl)

	if err != nil {
		errorModel.FailOnErr(err, "Failed Dial Rabbit Server...")
		return
	}
	RabbitChannel, err = conn.Channel()

	if err != nil {
		errorModel.FailOnErr(err, "Failed open Channel Rabbit Server...")
		return
	}
}

func connectRabbitMQ(url string) (*amqp091.Connection, error) {
	for {
		conn, err := amqp091.Dial(url)
		if err == nil {
			return conn, nil
		}
		log.Printf("Failed to connect to RabbitMQ: %v. Retrying in 5 seconds...", err)
		time.Sleep(5 * time.Second)
	}
}
