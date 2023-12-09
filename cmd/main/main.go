package main

import (
	"log"
	"net/http"
	"os"
	"rest-api-tutorial/internal/user"
	"time"

	"github.com/julienschmidt/httprouter"
)

var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	// Write log
	var file, err = os.OpenFile("/log/all.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	infoLog = log.New(file, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(file, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	infoLog.Println("Create router")
	router := httprouter.New()

	infoLog.Println("Register user handler")
	handler := user.NewHandler()
	handler.Register(router)
	start(router)
}

func start(router *httprouter.Router) {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     errorLog,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	infoLog.Println("The server is running and listening on port: 8080")
	errorLog.Fatal(server.ListenAndServe())
}
