package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"rest-api-tutorial/internal/config"
	"rest-api-tutorial/internal/user"
	"rest-api-tutorial/pkg/logging"
	"time"
)

const (
	cfgPath = "config.yaml"
	lgrPath = "logs/all.log"
)

func main() {
	// Logger
	lgr, err := logging.NewLogging(lgrPath)

	// Config
	cfg, err := config.NewConfig(cfgPath)
	if err != nil {
		lgr.ErrorLog.Fatal(err)
	}

	lgr.InfoLog.Println("Create router")
	router := httprouter.New()

	lgr.InfoLog.Println("Register user handler")
	handler := user.NewHandler()
	handler.Register(router)
	start(router, cfg, lgr)

}

func start(router *httprouter.Router, cfg *config.Config, lgr *logging.Logging) {
	server := &http.Server{
		Addr:         cfg.Server.Host + ":" + cfg.Server.Port,
		Handler:      router,
		ErrorLog:     lgr.ErrorLog,
		ReadTimeout:  cfg.Server.Timeout.Read * time.Second,
		WriteTimeout: cfg.Server.Timeout.Write * time.Second,
		IdleTimeout:  cfg.Server.Timeout.Idle * time.Second,
	}
	lgr.InfoLog.Printf("The server is running and listening on port: %s", cfg.Server.Port)
	log.Fatal(server.ListenAndServe())
}
