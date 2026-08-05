package rabbit

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewPublisher() (*Publisher, error) {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}
	err = ch.ExchangeDeclare(
		"products",
		"direct",
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
	productCreatedQueue, err := ch.QueueDeclare(
		"product-created",
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
	productDeletedQueue, err := ch.QueueDeclare(
		"product-deleted",
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

	if err := ch.QueueBind(
		productCreatedQueue.Name,
		"product.created",
		"products",
		false,
		nil,
	); err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}
	if err := ch.QueueBind(
		productDeletedQueue.Name,
		"product.deleted",
		"products",
		false,
		nil,
	); err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}

	return &Publisher{
		conn: conn,
		ch:   ch,
	}, nil
}

func (p *Publisher) Close() error {
	if err := p.ch.Close(); err != nil {
		return err
	}

	return p.conn.Close()
}

func (p *Publisher) Publish(routingKey string, body []byte) error {
	err := p.ch.Publish(
		"products",
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
