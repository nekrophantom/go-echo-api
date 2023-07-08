package main

import (
	"crud-simple-api/config"
	"crud-simple-api/db"
	"crud-simple-api/helper"
	"crud-simple-api/routes"
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	
	config.LoadConfig()

	err:= db.Init()
	helper.PanicIfError(err)

	err = db.Migrate(db.DB)
	if err != nil{
		log.Fatalf("Failed to migrate the database: %v", err)
	}

	e := echo.New()

	routes.SetupRoutes(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}