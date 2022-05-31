package models
import "gorm.io/gorm"

type Order struct {
	gorm.Model
	name string
	cardNo string
	cvc string
	date string
}