package main

import (
	"kv-mem/internal"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := internal.Bootstrap(port); err != nil {
		log.Fatal(err)
	}
}
