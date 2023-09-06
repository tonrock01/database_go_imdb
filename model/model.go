package model

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	ImdbID      string  `gorm:"unique";json:"imdbID"`
	Title       string  `json:"title"`
	Year        int     `json:"year"`
	Rating      float32 `json:"rating"`
	IsSuperHero bool    `json:"isSuperHero"`
}
