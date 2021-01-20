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
	urlEncode := base62.StdEncoding.EncodeToString([]byte(url))
	config.InitializeDatabase().QueryRow("INSERT INTO `apis` (`default_url`, `rewrite_url`) VALUES ('"+ url +"', 'http://localhost:5002/"+ urlEncode +"')")
	
}

func SearchUrl(url string) *Api {
	var api Api
	err := config.InitializeDatabase().QueryRow("SELECT `id`, `default_url`, `rewrite_url` FROM `apis` WHERE `rewrite_url` = ?", url).Scan(&api.Id, &api.DefaultUrl, &api.RewriteUrl)

	if err != nil {
		log.Fatal(err)
	}
  
	return &api
}