package main

import (
	"errors"
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	_, err := ReadContents("http://example.com/not_found")
	if err != nil {
		switch e := err.(type) {
		case *HTTPError:
			fmt.Printf("%T: %v\n", e, e.Error())
		default:
			panic(e)
		}
	}
}

type HTTPError struct {
	StatusCode int
	URL        string
}

func (he *HTTPError) Error() string {
	return fmt.Sprintf("http status code = %d, url = %s", he.StatusCode, he.URL)
}

func ReadContents(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, &HTTPError{
			StatusCode: resp.StatusCode,
			URL:        url,
		}
	}

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return byteArray, nil
}
