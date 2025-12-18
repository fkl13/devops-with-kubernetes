package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT must be set")
	}

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
	}

	log.Printf("Server started in port %s\n", port)

	err := srv.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		log.Println("Server closed")
	} else if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
