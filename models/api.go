package models

import (
	"log"
	"../config"
	"github.com/lytics/base62"
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

	urlEncode := base62.StdEncoding.EncodeToString([]byte(url))
	err := config.InitializeDatabase().QueryRow("INSERT INTO apis (`default_url`, `rewrite_url`) VALUES ('"+ url +"', 'http://localhost:5002/"+ urlEncode +"')")
  
	if err == nil {
	  	log.Fatal(err)
	}
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