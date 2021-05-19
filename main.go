package main

import (
	userhandler "babyFood/handlers"
	"babyFood/pkg/db"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	db.InitializeDBConnection()

	e := echo.New()

	e.GET("", userhandler.GetUsers)
	e.GET("/:id", userhandler.GetUser)

	e.Start(":8080")
	fmt.Printf("Running...")

}
