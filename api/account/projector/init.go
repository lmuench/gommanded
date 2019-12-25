package projector

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/lmuench/gommanded/typ"

	"github.com/lmuench/gommanded/api/account/event"

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
		for msg := range msgs {
			handleMessage(msg)
		}
	}()

	log.Println("Account projector waiting for events")
}

func handleMessage(msg amqp.Delivery) {
	log.Printf("Account Projector received message %s\n", msg.Body)
	var accountOpened event.AccountOpened
	err := json.Unmarshal(msg.Body, &accountOpened)
	if err != nil {
		log.Println("Account Projector failed to parse message as `accountOpened` event")
		return
	}
	log.Println("Account projector parsed message as `accountOpened` event")

	newAccount := typ.Account{
		UUID:      accountOpened.AccountUUID,
		CreatedAt: time.Now(),
		Balance:   accountOpened.InitialBalance,
		Closed:    false,
	}
	key := datastore.NameKey("Account", newAccount.UUID, nil)
	_, err = client.Put(ctx, key, &newAccount)
	if err == nil {
		log.Println("Account Projector persisted new `Account`:", newAccount)
	} else {
		log.Println("Account Projector failed to persist new `Account`:", err)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
