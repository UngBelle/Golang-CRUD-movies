package controllers

import (
	"crud-golang/initializers"
	"crud-golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MoviesCreate(c *gin.Context) {
	// Get data
	var movie struct {
		Name     string
		Producer string
		Year     string
	}

	c.Bind(&movie)

	// Create a movie
	m := models.Movie{Name: movie.Name, Producer: movie.Producer, Year: movie.Year}

	result := initializers.DB.Create(&m)
	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return
	c.JSON(http.StatusOK, gin.H{
		"movies": m,
	})
}

func MoviesGetAll(c *gin.Context) {
	// Get movies
	var movies []models.Movie
	initializers.DB.Find(&movies)

	// Respond
	c.JSON(200, gin.H{
		"movies": movies,
	})
}

func MovieGetOne(c *gin.Context) {
	// Get ID off URL
	id := c.Param("id")

	// Get the movie
	var amovie models.Movie
	initializers.DB.First(&amovie, id)
	// SELECT * FROM movies WHERE id = ...

	// Respond
	c.JSON(200, gin.H{
		"movies": amovie,
	})
}

func MovieUpdate(c *gin.Context) {
	// Get ID
	id := c.Param("id")

	// Get a movie
	var movie struct {
		Name     string
		Producer string
		Year     string
	}
	c.Bind(&movie)

	// Find the movie
	var amovie models.Movie
	initializers.DB.First(&amovie, id)

	// Update it
	initializers.DB.Model(&amovie).Updates(models.Movie{
		Name:     movie.Name,
		Producer: movie.Producer,
		Year:     movie.Year,
	})

	// Respond
	c.JSON(200, gin.H{
		"movies": amovie,
	})
}

func MovieDelete(c *gin.Context) {
	// Get ID
	id := c.Param("id")

	// Delete a movie
	initializers.DB.Delete(&models.Movie{}, id)

	// Respond
	c.Status(200)
}
