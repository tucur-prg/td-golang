package main

import (
	"fmt"
	"net/http"
	"embed"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed web
var app embed.FS

func main() {
	fmt.Println("--- start ---")

	go runStaticServer()
	go runApiServer()

	select {}
}

func runStaticServer() {
	e := echo.New()

	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       "web",
		HTML5:      true,
		Filesystem: http.FS(app),
	}))

	e.Logger.Fatal(e.Start(":9090"))
}

func runApiServer() {
	e := echo.New()

	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/json", func(c echo.Context) error {
		j := map[string]string{
			"hoge": "fuga",
			"key": "value",
		}
		return c.JSON(http.StatusOK, j)
	})

	e.Logger.Fatal(e.Start(":9091"))
}
