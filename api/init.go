package api

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
	"github.com/streadway/amqp"
)

var (
	conn   *amqp.Connection
	ch     *amqp.Channel
	q      amqp.Queue
	ctx    context.Context
	client *datastore.Client
)

func Init(context context.Context, dsClient *datastore.Client) {
	ctx, client = context, dsClient
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
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
