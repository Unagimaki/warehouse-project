package rabbit

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Sender struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewSender() (*Sender, error) {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	err = SetupProducts(ch)

	if err != nil {
		ch.Close()
		return nil, err
	}

	return &Sender{
		conn: conn,
		ch:   ch,
	}, nil
}

func (p *Sender) Close() error {
	if err := p.ch.Close(); err != nil {
		return err
	}

	return p.conn.Close()
}

func (p *Sender) Send(routingKey string, body []byte) error {
	fmt.Println("sending to rabbit:", routingKey, string(body))

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
