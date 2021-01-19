package controllers

import (
	"net/http"
	"encoding/json"
	//"../models"
)

func NewUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode("OK")
}

func SearchUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode("OK")
}