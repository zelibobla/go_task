package queue

import (
	"github.com/streadway/amqp"
)

type Queue interface {
	SendMessage(message string) error
	ReceiveMessages() (<-chan amqp.Delivery, error) // Add this method
}