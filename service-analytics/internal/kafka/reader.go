package kafka

import "github.com/segmentio/kafka-go"

func NewReader() *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "warehouse.public.products",
		GroupID: "analytics",
	})
}
