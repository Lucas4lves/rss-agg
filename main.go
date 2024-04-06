package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	authorName := "Lucas"

	fmt.Printf("RSS Aggregator | by: %s\n", authorName)

	godotenv.Load("./.env")

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Error while reading PORT env variable!\n")
	}
	fmt.Printf("PORT: %s\n", port)

}
