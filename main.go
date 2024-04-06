package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
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

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/health", handleReadiness)

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	fmt.Printf("Server listening on PORT: %s\n", port)
	err := http.ListenAndServe(server.Addr, server.Handler)
	if err != nil {
		log.Fatal("Error on initializing the server")
	}

}
