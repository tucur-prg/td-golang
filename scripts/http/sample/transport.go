package sample

import (
	"fmt"
	"time"
	"net/http"
)

type Transport struct {
	base http.RoundTripper
}

func NewTransport(base http.RoundTripper) http.RoundTripper {
	return &Transport{
		base: base,
	}
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	fmt.Println("RoundTrip: start")

	req.Header.Set("X-Sample-Request", "xxxxx")

	t1 := time.Now()

	res, err := t.base.RoundTrip(req)

	t2 := time.Now()
	fmt.Println("RoundTrip: end")
	fmt.Println(t2.Sub(t1))

	return res, err
}