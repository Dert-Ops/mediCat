package rabbitmq

import (
	"github.com/streadway/amqp"
	"gitlab.com/dert-ops/mediCat/mediCat-Dev.git/cmd/config"
)

var conn *amqp.Connection
var ch *amqp.Channel

func InitRabbitMQ() {
	var err error
	conn, err = amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		config.LogrusLogger.Fatalf("Failed Connect RABBİT MınaKoyim %s", err)
	}
	ch, err = conn.Channel()
	if err != nil {
		config.LogrusLogger.Fatalf("Failed to open Channel %s", err)
	}
}

func CreateQueue(queueName string) {
	_, err := ch.QueueDeclare(
		queueName, // Kuyruk ismi
		false,     // Durable
		false,     // Delete when unused
		false,     // Exclusive
		false,     // No-wait
		nil,       // Arguments
	)
	if err != nil {
		config.LogrusLogger.Fatalf("Failed to declare a queue: %v", err)
	}
}

// Mesaj gönder
func PublishMessage(queueName string, message string) {
	err := ch.Publish(
		"",        // Exchange
		queueName, // Routing key (kuyruk adı)
		false,     // Mandatory
		false,     // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		config.LogrusLogger.Fatalf("Failed to publish a message: %v", err)
	}
}
