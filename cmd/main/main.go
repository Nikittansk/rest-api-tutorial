package main

import (
	"log"
	"net/http"
	"rest-api-tutorial/internal/user"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	log.Println("Create router")
	router := httprouter.New()

	log.Println("Register user handler")
	handler := user.NewHandler()
	handler.Register(router)
	start(router)
}

func start(router *httprouter.Router) {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	log.Println("The server is running and listening on port: 8080")
	log.Fatal(server.ListenAndServe())
}
