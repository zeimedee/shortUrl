package models

import "gorm.io/gorm"

type InitialUrl struct {
	Url string
}

type ShortUrl struct {
	Url string
}

type Urls struct {
	gorm.Model

	InitialUrl string `json:"initialurl"`
	ShortUrl   string `json:"shorturl"`
}
