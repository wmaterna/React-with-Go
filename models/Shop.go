package models
import "gorm.io/gorm"

type Shop struct {
	gorm.Model
	Name string
	Address string
}