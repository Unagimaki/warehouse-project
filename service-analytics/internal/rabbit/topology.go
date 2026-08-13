package rabbit

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func SetupProducts(ch *amqp.Channel) error {
	err := ch.ExchangeDeclare(
		"products",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
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
		return err
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
		return err
	}
	err = ch.QueueBind(
		productCreatedQueue.Name,
		"product.created",
		"products",
		false,
		nil,
	)
	if err != nil {
		return err
	}
	err = ch.QueueBind(
		productDeletedQueue.Name,
		"product.deleted",
		"products",
		false,
		nil,
	)
	if err != nil {
		return err
	}
	return nil
}
