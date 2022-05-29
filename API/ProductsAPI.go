package API

import (
	"goShopGORM/db"
	"goShopGORM/models"
	"net/http"
	"github.com/labstack/echo"
)

func GetProducts(c echo.Context) error {
	db := db.InitDB()
	prod := []models.Product{}
	db.Find(&prod)
	return c.JSON(http.StatusOK, prod)
}

func AddProducts(e echo.Context) error {
	db := db.InitDB()
	prod := new(models.Product)
	if err := e.Bind(prod);
	err != nil {
		return err
	}
	db.Create(&prod)
	return e.String(http.StatusOK, "DB element added")
}