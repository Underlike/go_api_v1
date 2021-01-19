package main

import (
	"fmt"
	"log"
	"time"
	"net/http"
	"./controllers"
	"./config"
	"github.com/gorilla/mux"
)

const (
	InfoColor ="\033[1;34m%s\033[0m"
)

func main() {
	config.InitializeDatabase()
	initializeRouter()
}

func initializeRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/v1/api/new", controllers.NewUrl)
	router.HandleFunc("/v1/api/{url}", controllers.SearchUrl)

	srv := &http.Server{
        Handler: router,
        Addr: "127.0.0.1:5002",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

	fmt.Printf(InfoColor, "Server start on http://localhost:5002 \n")
	log.Fatal(srv.ListenAndServe())
}