package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"github.com/gorilla/mux"
)

func NewUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	url := r.FormValue("url")

	var api models.Api

	fmt.Printf("Insert into database down - OK \n")
	json.NewEncoder(w).Encode(models.NewUrl(&api, url))
}

func RedirectUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	api := models.SearchUrl(vars["key"])

	http.Redirect(w, r, api.Url, 302)
}
