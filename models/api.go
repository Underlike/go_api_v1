package models

import (
	"fmt"
	"log"
	"../config"
	//"github.com/kare/base62"
)

type Api struct {
	Id int `json:"id"`
	DefaultUrl string `json:"default_url"`
	RewriteUrl string `json:"rewrite_url"`
}

type Apis []Api

func NewUrl(api *Api, url string) {
	if api == nil {
	  	log.Fatal(api)
	}
  
	err := config.InitializeDatabase().QueryRow("INSERT INTO apis (`default_url`, `rewrite_url`) VALUES ('"+ url +"', '"+ url +"')")
  
	if err == nil {
	  	log.Fatal(err)
	}

	fmt.Printf("Insert into database down - OK \n")
}

func SearchUrl(url string) *Api {
	var api Api
  
	row := config.InitializeDatabase().QueryRow("SELECT * FROM apis WHERE rewrite_url = '" + url + "';")
	err := row.Scan(&api.Id, &api.DefaultUrl, &api.RewriteUrl)
  
	if err != nil {
	  	log.Fatal(err)
	}
  
	return &api
}