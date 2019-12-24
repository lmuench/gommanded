package projector

import (
	"log"

	"github.com/streadway/amqp"
)

var (
	conn *amqp.Connection
	ch   *amqp.Channel
	q    amqp.Queue
)

func Init() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	// TODO: call `conn.Close()` somewhere for graceful shutdown
	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")
	// TODO: call `ch.Close()` somewhere for graceful shutdown
	q, err = ch.QueueDeclare(
		"accounts", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			log.Printf("Account projector received a message: %s", d.Body)
		}
	}()

	log.Println("Account projector waiting for messages")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
