package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Name     string
	Producer string
	Year     string
}
