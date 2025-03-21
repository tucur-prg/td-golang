package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"embed"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"main/src"
)

type Example struct {
	Foo string `query:"foo" validate:"required"`
	Bar string `query:"bar" validate:"omitempty,print"`
}

//go:embed web
var app embed.FS

func main() {
	fmt.Println("--- start ---")

	e := echo.New()
	e.Validator = my.NewValidator()

	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(my.Middleware2)
	e.Use(my.Middleware1)

	g := e.Group("/test/:hoge")
	{
		g.GET("/1", func(c echo.Context) error {
			fmt.Println("group 1", c.Param("hoge"))
			return c.NoContent(http.StatusOK)
		})

		g.Use(middleware.StaticWithConfig(middleware.StaticConfig{
			Root:       "web",
			HTML5:      true,
			Filesystem: http.FS(app),
		}))
	
	}

	{
		h := my.NewHello("a")
		e.GET("/1", h.Call)
	}
	{
		h := my.NewHello("b")
		e.GET("/2", h.Call)
	}
	{
		h := my.NewHello("c")
		e.GET("/3/:id", h.Call, my.Middleware3, my.Middleware4)
	}
	{
		h := my.NewHello("e")
		e.GET("/3/me", h.Call, my.Middleware5)
	}
	{
		e.GET("/v", func(c echo.Context) error {
			obj := new(Example)
			if err := c.Bind(obj); err != nil {
				fmt.Println("bind error", err.Error())
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			if err := c.Validate(obj); err != nil {
				fmt.Println("validate error", err.Error())
				return err
			}

			fmt.Println("Example: ", obj)

			return c.String(http.StatusOK, "OK\n")
		})
	}
	{
		e.GET("/json", func(c echo.Context) error {
			j := map[string]string{
				"hoge": "fuga",
				"key": "value",
			}
			return c.JSON(http.StatusOK, j)
		})
		e.GET("/string", func(c echo.Context) error {
			j := map[string]string{
				"hoge": "fuga",
				"key": "value",
			}
			s, _ := json.Marshal(j)
			c.Response().Header().Set(echo.HeaderContentType, "application/json")
			return c.String(http.StatusOK, string(s))
		})
		e.GET("/html", func(c echo.Context) error {
			h := `<!DOCTYPE html>
<html>
<body>
<div>Hello, World!!</div>
</body>
</html>`
			return c.HTML(http.StatusOK, h)
		})
		e.GET("/redirect", func(c echo.Context) error {
			return c.Redirect(http.StatusFound, "/1")
		})
	}
	{
		e.GET("/users/:id", func(c echo.Context) error {
			fmt.Println("id:", c.Param("id"))
			fmt.Println(c.Echo().Reverse("users", "hoge"))
			return c.String(http.StatusOK, "ok!")
		}).Name = "users"
	}

	e.Logger.Fatal(e.Start(":9090"))
}
