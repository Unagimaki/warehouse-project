package main

import (
	"log"
	"service-werehouse/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
