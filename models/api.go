package models

import (
	"../config"
	"math/rand"
	"time"
	"github.com/kare/base62"
)

type Api struct {
	Id int `json:"id"`
	DefaultUrl string `json:"default_url"`
	RewriteUrl string `json:"rewrite_url"`
}

type Apis []Api

func NewUrl(api *Api, url string) *Api {
	var apiSend Api
	rand.Seed(time.Now().Unix())
	encodage := rand.Int63n(1566541545632156)
	urlEncode := "http://localhost:5002/" + base62.Encode(encodage)

	config.InitializeDatabase().QueryRow("INSERT INTO `apis` (`default_url`, `rewrite_url`) VALUES ('"+ url +"', '"+ urlEncode +"');")
	config.InitializeDatabase().QueryRow("SELECT `id`, `default_url`, `rewrite_url` FROM `apis` WHERE `rewrite_url` = ?", urlEncode).Scan(&api.Id, &api.DefaultUrl, &api.RewriteUrl)

	return &apiSend
}

func SearchUrl(url string) *Api {
	var api Api
	config.InitializeDatabase().QueryRow("SELECT `id`, `default_url`, `rewrite_url` FROM `apis` WHERE `rewrite_url` = ?", url).Scan(&api.Id, &api.DefaultUrl, &api.RewriteUrl)
  
	return &api
}