package API

import (
	"goShopGORM/db"
	"goShopGORM/models"
	"net/http"
	"github.com/labstack/echo"
)

func AddOrder(e echo.Context) error {
	db := db.InitDB()
	order := new(models.Order)
	if err := e.Bind(order);
	err != nil {
		return err
	}
	db.Create(&order)
	return e.String(http.StatusOK, "DB element added")
}