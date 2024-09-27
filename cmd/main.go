package main

import (
	"LinkShortener/internal/controller"
	"LinkShortener/internal/db/postgres"
	"LinkShortener/internal/db/redis"
	"LinkShortener/internal/router"
	"log"
	"net/http"
)

var HOST = "127.0.0.1"

func main() {
	err := postgres.Connect(HOST, "ls_admin", "Pa$w0rd", "linkShortenerDB", 5433)
	if err != nil {
		log.Fatal("Can't connect to db. Err:", err)
	}

	redis.Connect(HOST, 6380, "Pa$$w0rd", 0)

	handler := router.NewRouter(controller.NewController())

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	errors := make(chan error, 1)
	go func() {
		log.Println("Server Started")
		errors <- server.ListenAndServe()
	}()

	if err := <-errors; err != nil {
		log.Fatalf("Server stopped. Error: %v", err)
	}
}
