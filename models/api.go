package models

import (
	"log"

	"../config"
	"github.com/kare/base62"
)

type Api struct {
	Id  int
	Url string
}

type Apis []Api

func NewUrl(api *Api, url string) *string {
	config.InitializeDatabase().Order("id desc").Find(&api)

	urlEncode := "http://localhost:5002/" + base62.Encode(int64(api.Id))
	ap := Api{Url: url}

	config.InitializeDatabase().Create(&ap)

	return &urlEncode
}

func SearchUrl(key string) *Api {
	var api Api

	id, err := base62.Decode(key)

	if err != nil {
		log.Println(err)
	}

	config.InitializeDatabase().Table("apis").Select("url").Where("id = ?", id).Scan(&api)

	return &api
}
