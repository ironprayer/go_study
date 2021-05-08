package main

import (
	"github.com/labstack/echo/v4"
)

func handleHome(c echo.Context) error {
	return c.File("index.html")
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.Logger.Fatal(e.Start(":1111"))
}
