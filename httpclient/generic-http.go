package httpclient

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func doHTTPGet(url string) (io.ReadCloser, error) {

	log.Printf("get request for : %s", url)
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create new request %v", err)
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to perform get request %v", err)
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("here api returned status code %s", response.Status)
	}
	log.Println("get method success")
	return response.Body, nil
}
