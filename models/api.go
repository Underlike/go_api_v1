package models

type Api struct {
	Id int `json:"Id"`
	DefaultUrl string `json:"DefaultUrl"`
	RewriteUrl string `json:"RewriteUrl"`
}

type Apis []Api