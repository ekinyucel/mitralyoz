package http

import (
	"fmt"
	"net/http"
	"time"
)

// SendHTTPRequest is used for sending a http request to given target
func SendHTTPRequest(url string) {
	start := time.Now()
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error calling: %v\n", url)
		return
	}

	elapsedTime := time.Since(start)

	fmt.Printf("Successful call to: %v total elapsed time %v\n", url, elapsedTime)
	fmt.Println()
	fmt.Println("status code", response.StatusCode)
}
