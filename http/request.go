package http

import (
	"fmt"
	"net/http"
)

// SendHTTPRequest is used for sending a http request to given target
func SendHTTPRequest(url string) {
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error calling: %v\n", url)
		return
	}
	fmt.Printf("Successful call to: %v\n", url)
	fmt.Println()
	fmt.Println("status code", response.StatusCode)
}
