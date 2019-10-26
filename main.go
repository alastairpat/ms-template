package main

import (
	"github.com/alastairpat/ms-template/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handlers.GetRootIndex)
	e.GET("/status", handlers.GetStatusIndex)

	e.Logger.Fatal(e.Start(":8080"))
}
