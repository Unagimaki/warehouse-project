package main

import (
	"encoding/json"
	"fmt"
	"log"

	"service-analytics/internal/consumer"
	"service-analytics/internal/kafka"
	"service-analytics/internal/rabbit"
)

type Event struct {
	Payload Payload `json:"payload"`
}

type Payload struct {
	Before *Product `json:"before"`
	After  *Product `json:"after"`
	Op     string   `json:"op"`
}

type Product struct {
	ID          int    `json:"id"`
	ItemNumber  string `json:"item_number"`
	Type        string `json:"type"`
	Brand       string `json:"brand"`
	Gender      string `json:"gender"`
	VolumeML    int    `json:"volume_ml"`
	Description string `json:"description"`
	Barcode     string `json:"barcode"`
	CreatedAt   int64  `json:"created_at"`
}

func main() {
	reader := kafka.NewReader()
	sender, err := rabbit.NewSender()
	if err != nil {
		log.Fatal(err)
	}
	rabbitConsumer, err := rabbit.NewConsumer()
	if err != nil {
		log.Fatal(err)
	}

	deliveries, err := rabbitConsumer.Consume()
	if err != nil {
		log.Fatal(err)
	}

	for msg := range deliveries {
		event, err := ParseEvent(msg.Body)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if err := msg.Ack(false); err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(event.Payload.Op)
		fmt.Println(event.Payload.After)
	}
	defer reader.Close()
	defer sender.Close()

	consumer.StartKafkaConsumer(reader, sender)
}

func ParseEvent(data []byte) (Event, error) {
	var event Event

	if err := json.Unmarshal(data, &event); err != nil {
		return Event{}, err
	}

	return event, nil
}
