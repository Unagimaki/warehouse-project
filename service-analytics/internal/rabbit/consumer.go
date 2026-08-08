package rabbit

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewConsumer() (*Consumer, error) {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}
	return &Consumer{
		conn: conn,
		ch:   ch,
	}, nil
}

func (c *Consumer) Consume() (<-chan amqp.Delivery, error) {
	deliveries, err := c.ch.Consume(
		"product-created",
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return deliveries, nil
}
