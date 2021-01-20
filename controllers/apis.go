package controllers

import (
	"fmt"
	"net/http"
	"../models"
	"encoding/json"
	"github.com/gorilla/mux"
)

func NewUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	url := r.FormValue("url")

	var api models.Api
	models.NewUrl(&api, url)
	json.NewEncoder(w).Encode("success")
	fmt.Printf("Insert into database down - OK \n")
}

func SearchUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode("OK")

	vars := mux.Vars(r)
	url := "http://localhost:5002/" + vars["key"]

	api := models.SearchUrl(url)
	fmt.Printf(api.DefaultUrl)
	http.Redirect(w, r, api.DefaultUrl, 302)
}