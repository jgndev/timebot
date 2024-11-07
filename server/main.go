// main.go
package main

import (
	"github.com/jgndev/timebot/internal/application"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	//e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
	//	return func(c echo.Context) error {
	//		c.Response().Header().Set("Cache-Control", "public, max-age=86400")
	//		return next(c)
	//	}
	//})

	// Static assets
	e.Static("/public", "public")
	e.File("/favicon.ico", "public/img/favicon.ico")
	e.File("/robots.txt", "public/txt/robots.txt")

	// Routes
	e.GET("/", application.Home)
	e.GET("/times", application.TimeUpdate)

	// Start app
	e.Logger.Fatal(e.Start(":8080"))
}
