package API

import (
	"goShopGORM/db"
	"goShopGORM/models"
	"net/http"
	"github.com/labstack/echo"
)

func GetUsers(c echo.Context) error {
	db := db.InitDB()
	users := []models.User{}
	db.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func AddUser(e echo.Context) error {
	db := db.InitDB()
	user := new(models.User)
	if err := e.Bind(user);
	err != nil {
		return err
	}
	db.Create(&user)
	return e.String(http.StatusOK, "DB element added")
}