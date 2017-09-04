package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.Static("/", "../frontend/dist")

	e.Logger.Fatal(e.Start(":8080"))
}
