package models

import (
	"../config"
	"github.com/kare/base62"
)

type Api struct {
	Id         int
	DefaultUrl string
	RewriteUrl string
}

type Apis []Api

func NewUrl(api *Api, url string) **Api {
	config.InitializeDatabase().Order("id desc").Find(&api)

	urlEncode := "http://localhost:5002/" + base62.Encode(int64(api.Id))
	ap := Api{DefaultUrl: url, RewriteUrl: urlEncode}

	config.InitializeDatabase().Create(&ap)
	config.InitializeDatabase().Raw("SELECT `id`, `default_url`, `rewrite_url` FROM `apis` WHERE `rewrite_url` = ?", urlEncode).Scan(&api)

	return &api
}

func SearchUrl(url string) *Api {
	var api Api
	config.InitializeDatabase().Raw("SELECT `id`, `default_url`, `rewrite_url` FROM `apis` WHERE `rewrite_url` = ?", url).Scan(&api)

	return &api
}
