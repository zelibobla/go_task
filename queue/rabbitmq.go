package queue

import (
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewRabbitMQ(url, queueName string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}

	return &RabbitMQ{
		conn:    conn,
		channel: ch,
		queue:   q,
	}, nil
}

func (r *RabbitMQ) SendMessage(message string) error {
	return r.channel.Publish(
		"",
		r.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
}

func (r *RabbitMQ) ReceiveMessages() (<-chan amqp.Delivery, error) {
	msgs, err := r.channel.Consume(
		r.queue.Name,
		"",
		true,  // auto-acknowledge
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,
	)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}

func (r *RabbitMQ) Close() {
	r.channel.Close()
	r.conn.Close()
}
