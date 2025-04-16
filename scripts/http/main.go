package main

import (
	"fmt"
	"net/http"

	"http/sample"
)

func main() {
	client := &http.Client{
		Transport: sample.NewTransport(
			sample.NewTransport(http.DefaultTransport),
		),
	}

	req, _ := http.NewRequest(http.MethodGet, "https://example.com", nil)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T: %+v\n", req, req.Header)
	fmt.Printf("%T: %+v\n", resp, resp.Header)
}