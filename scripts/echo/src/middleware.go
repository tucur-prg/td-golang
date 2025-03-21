package my

import (
	"crypto/rand"
	"errors"
	"fmt"

	"github.com/labstack/echo/v4"
)

func MakeRandomStr(digit uint32) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 乱数を生成
	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error...")
	}

	// letters からランダムに取り出して文字列を生成
	var result string
	for _, v := range b {
        // index が letters の長さに収まるように調整
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}

func Middleware1(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
		fmt.Println("-- Middleware 1 Call")

		c.Set("counter", 0) // 初期化

		return next(c)
    }
}
func Middleware2(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
		fmt.Println("-- Middleware 2 Call")

        endpoint := "UNKNOWN"
        for _, r := range c.Echo().Routes() {
			// https://pkg.go.dev/github.com/labstack/echo/v4#Route
			// r = echo.Route
            if r.Method == c.Request().Method && r.Path == c.Path() {
                endpoint = r.Name
            }
        }

		fmt.Println("Route Name:", endpoint)

		fmt.Println("Scheme: ", c.Scheme())
		fmt.Println("RealIP: ", c.RealIP())
		fmt.Println("Path: ", c.Path())
		fmt.Println("QueryString: ", c.QueryString())

		fmt.Println("Request: ", c.Request())

		err := next(c)

		fmt.Println("Response: ", c.Response())

		return err
    }
}
func Middleware3(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
		fmt.Println("-- Middleware 3 Call")

		return next(c)
    }
}
func Middleware4(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
		fmt.Println("-- Middleware 4 Call")

		return next(c)
    }
}
func Middleware5(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
		fmt.Println("-- Middleware 5 Call")

//		fmt.Println(c.Param())

		fmt.Println("rand:", c.Get("rand"))
		r, _ := MakeRandomStr(10)
		c.Set("rand", r)

		c.SetParamNames("id")
		c.SetParamValues("self")

		fmt.Println(c.ParamNames())
		fmt.Println(c.ParamValues())

		err := next(c)

		c.Set("rand2", r)

		return err
    }
}
