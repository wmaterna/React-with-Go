package models
import "gorm.io/gorm"

type Order struct {
	gorm.Model
	// DateOfOrder time.Time
	UserID uint
}