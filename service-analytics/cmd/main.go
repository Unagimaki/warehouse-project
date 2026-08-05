package main

import (
	"context"
	"fmt"

	"service-analytics/internal/kafka"
)

func main() {
	reader := kafka.NewReader()

	defer reader.Close()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(string(msg.Value))
	}
}
