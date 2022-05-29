package routes

import(
	"goShopGORM/API"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/users", API.GetUsers)
	e.POST("/users", API.AddUser)
	e.GET("/products", API.GetProducts)
	e.POST("/products", API.AddProducts)
	return e
}