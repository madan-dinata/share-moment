package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World Echo Golang")
	})
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":5000"))
}