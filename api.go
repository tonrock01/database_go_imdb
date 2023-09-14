package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	r := gin.Default()
	r.GET("/movies", func(c *gin.Context) {
		var movies []model.Movie

		db.Find(&movies)
		c.JSON(200, movies)
	})
	r.GET("/movies/:imdbID", func(c *gin.Context) {
		imdbID := c.Param("imdbID")
		var movie model.Movie
		db.Where("imdb_id = ?", imdbID).First(&movie)
		c.JSON(http.StatusOK, movie)
	})
	r.GET("/movies/years/:year", func(c *gin.Context) {
		year := c.Param("year")
		var movie []model.Movie
		db.Where("year = ?", year).Find(&movie)
		c.JSON(http.StatusOK, movie)
	})
	r.POST("/movies", func(c *gin.Context) {
		var movie model.Movie
		if err := c.Bind(&movie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result := db.Create(&movie)
		c.JSON(http.StatusOK, gin.H{"RowsAffected": result.RowsAffected})
	})
	r.DELETE("/movies/:imdbID", func(c *gin.Context) {
		imdbID := c.Param("imdbID")
		var movie model.Movie
		db.Where("imdb_id = ?", imdbID).First(&movie)
		db.Delete(&movie)
		c.JSON(http.StatusOK, gin.H{"Delete Movie": movie.ImdbID})
	})
	r.PUT("/movies", func(c *gin.Context) {
		var movie model.Movie
		var updatedmovie model.Movie
		if err := c.Bind(&movie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Where("imdb_id = ?", movie.ImdbID).First(&updatedmovie)
		updatedmovie.Title = movie.Title
		updatedmovie.Year = movie.Year
		updatedmovie.Rating = movie.Rating
		updatedmovie.IsSuperHero = movie.IsSuperHero
		db.Save(updatedmovie)

		c.JSON(http.StatusOK, gin.H{"Update Movie": updatedmovie})
	})
	r.Run()
}
