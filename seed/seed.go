package main

import (
	"github.com/tonrock01/goimdbSQL/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root@tcp(127.0.0.1:3306)/goimdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.Movie{})

	// Create
	db.Create(&model.Movie{ImdbID: "tt4154796", Title: "Avengers: Endgame", Year: 2019, Rating: 8.4, IsSuperHero: true})
	db.Create(&model.Movie{ImdbID: "tt4154683", Title: "Black Panther", Year: 2018, Rating: 7.3, IsSuperHero: false})

}
