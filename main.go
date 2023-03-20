package main

import (
	"fmt"
	"todoapi/config"
	"todoapi/config/db"
	"todoapi/controller"
	"todoapi/model"
	"todoapi/router"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	err := godotenv.Load("local.env")
	if err != nil {
		fmt.Printf("failed to lead env file: %v \n", err)
	}

	cfg := config.NewConfig()

	sqliteDB, err := db.InitDb(cfg.ConnDb)
	if err != nil {
		panic("failed to connect to db")
	}
	sqliteDB.AutoMigrate(&model.Todo{})

	storage := controller.NewStorage(sqliteDB)

	e := echo.New()

	router.InitRouter(e, storage)

	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
