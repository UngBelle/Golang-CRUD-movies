package main

import (
	"crud-golang/controllers"
	"crud-golang/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/movies", controllers.MoviesCreate)
	r.GET("/movies", controllers.MoviesGetAll)
	r.GET("/movie/:id", controllers.MovieGetOne)
	r.PUT("/movie/:id", controllers.MovieUpdate)
	r.DELETE("/movie/:id", controllers.MovieDelete)

	r.Run()
}
