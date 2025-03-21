package my

import (
	"net/http"
	"fmt"
	"context"

	"github.com/labstack/echo/v4"
)

type Hello interface {
	Call(c echo.Context) error
}

type hello struct {
	counter int
	value string
	ctx context.Context
}

func NewHello(v string) Hello {
	return &hello{counter: 0, value: v, ctx: context.Background()}
}

func (h *hello) Call(c echo.Context) error {
	h.counter++

	ctx1 := context.WithValue(h.ctx, "key1", h.value)
	ctx2 := context.WithValue(ctx1, "key2", h.value)
	fmt.Println(ctx2)

	fmt.Println(c.Response().Header().Get(echo.HeaderXRequestID))

	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", h.ctx)
	fmt.Printf("%T\n", ctx2)

	fmt.Println(c.Get("counter"))

	fmt.Println("ID:", c.Param("id"))
	fmt.Println("rand:", c.Get("rand"))
	fmt.Println("rand2:", c.Get("rand2"))

	return c.String(http.StatusOK, fmt.Sprintf("Hello %d\n", h.counter))
}
