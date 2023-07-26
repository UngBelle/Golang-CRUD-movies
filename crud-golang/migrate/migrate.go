package main

import (
	"crud-golang/initializers"
	"crud-golang/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Movie{})
}
