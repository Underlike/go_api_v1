package controllers

import (
	"fmt"
	//"log"
	//"strconv"
	"net/http"
	"../models"
	"encoding/json"
	//"github.com/gorilla/mux"
	
)

func NewUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	url := "https://test.fr"

	var api models.Api
	models.NewUrl(&api, url)
	json.NewEncoder(w).Encode("success")
	fmt.Printf("Insert into database down - OK \n")
}

func SearchUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode("OK")

	/**
	 * Faire passer la variable en body
	 * Rediriger vers l'original
	 */

	api := models.SearchUrl("https://test.fr/")
	json.NewEncoder(w).Encode(api)
}