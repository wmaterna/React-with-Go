package db

import(
	"goShopGORM/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


var db *gorm.DB
// var err error

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"))
  if err != nil {
    panic("failed to connect database")
  }

  db.AutoMigrate(&models.User{})
  db.AutoMigrate(&models.Order{})
  db.AutoMigrate(&models.Product{})
  db.AutoMigrate(&models.Supplier{})
  db.AutoMigrate(&models.Shop{})

  return db
}

// func ConnectDatabase() *gorm.DB {
// 	return db
// }