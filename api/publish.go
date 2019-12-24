package api

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

func publish(v interface{}) {
	msg, err := json.Marshal(v)
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		},
	)
	failOnError(err, "Failed to publish a message")
}
