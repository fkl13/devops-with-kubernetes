package main

import (
	"log"
	"time"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New()

	for {
		log.Println(id)
		time.Sleep(5 * time.Second)
	}
}
