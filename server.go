package main

import (
  "goShopGORM/db"
  "goShopGORM/routes"
)



func main() {
  db.InitDB()
  e := routes.Init()
  e.Start(":8080")
}


  