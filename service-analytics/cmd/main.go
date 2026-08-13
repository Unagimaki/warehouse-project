package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"service-analytics/internal/consumer"
	"service-analytics/internal/domain"
	"service-analytics/internal/kafka"
	"service-analytics/internal/opensearch"
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
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	reader := kafka.NewReader()
	sender, err := rabbit.NewSender()
	if err != nil {
		return err
	}

	osClient, err := opensearch.NewClient()
	if err != nil {
		return err
	}

	rabbitConsumer, err := rabbit.NewConsumer()
	if err != nil {
		return err
	}

	deliveries, err := rabbitConsumer.Consume()
	if err != nil {
		return err
	}

	defer reader.Close()
	defer sender.Close()
	go consumer.StartKafkaConsumer(reader, sender)

	for msg := range deliveries {
		fmt.Println("received from rabbit:", string(msg.Body))
		event, err := ParseEvent(msg.Body)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if event.Payload.After == nil {
			continue
		}

		prod := toDomainProduct(*event.Payload.After)
		err = osClient.IndexProduct(context.Background(), prod)
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

	return nil
}

func toDomainProduct(p Product) domain.Product {
	return domain.Product{
		ID:          p.ID,
		ItemNumber:  p.ItemNumber,
		Type:        p.Type,
		Brand:       p.Brand,
		Gender:      p.Gender,
		VolumeML:    p.VolumeML,
		Description: p.Description,
		Barcode:     p.Barcode,
		CreatedAt:   p.CreatedAt,
	}
}

func ParseEvent(data []byte) (Event, error) {
	var event Event

	if err := json.Unmarshal(data, &event); err != nil {
		return Event{}, err
	}

	return event, nil
}
