package controllers

import (
	"net/http"
	"encoding/json"
	"../models"
)

func NewUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	/**
	 * Faire passer la variable en body
	 */

	var api models.Api
	models.NewUrl(&api, "https://test.fr")
	json.NewEncoder(w).Encode(api)
}

func SearchUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode("OK")

	/**
	 * Faire passer la variable en body
	 */

	api := models.SearchUrl("https://test.fr/")
	json.NewEncoder(w).Encode(api)
}