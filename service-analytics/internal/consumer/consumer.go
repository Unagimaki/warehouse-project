package consumer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"

	"service-analytics/internal/rabbit"
)

type Event struct {
	Payload Payload `json:"payload"`
}

type Payload struct {
	Op string `json:"op"`
}

func StartKafkaConsumer(reader *kafka.Reader, sender *rabbit.Sender) {
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println(err)
			continue
		}
		op, err := GetOperation(msg.Value)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(op)
		routingKey, err := getRoutingKey(op)
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = sender.Send(routingKey, msg.Value)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func GetOperation(data []byte) (string, error) {
	var event Event

	if err := json.Unmarshal(data, &event); err != nil {
		return "", err
	}

	return event.Payload.Op, nil
}
func getRoutingKey(op string) (string, error) {
	var routingKey string

	switch op {
	case "c":
		routingKey = "product.created"
	case "u":
		routingKey = "product.updated"
	case "d":
		routingKey = "product.deleted"
	default:
		return "", fmt.Errorf("unknown operation: %s", op)
	}
	return routingKey, nil
}
