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
	router := mux.NewRouter().StrictSlash(true)
	
	router.Methods("POST").Path("/v1/api/new").Name("Create").HandlerFunc(controllers.NewUrl)
	router.Methods("GET").Path("/{key}").Name("Redirect").HandlerFunc(controllers.RedirectUrl)

	srv := &http.Server{
        Handler: router,
        Addr: "127.0.0.1:5002",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

	fmt.Printf(InfoColor, "Server start on http://localhost:5002 \n")
	log.Fatal(srv.ListenAndServe())
}