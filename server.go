package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", health)

	e.Logger.Fatal(e.Start(":8081"))
}

func health(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
