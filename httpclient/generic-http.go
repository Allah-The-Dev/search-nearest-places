package httpclient

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func doHTTPGet(url string) (*io.ReadCloser, error) {

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("here geocode api returned status code %s", response.Status)
	}

	return &response.Body, nil
}
