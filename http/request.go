package http

import (
	"log"
	"net/http"
	"time"
)

// Result is holding the information about the executed HTTP result
type Result struct {
	ResultID    int
	URL         string
	ElapsedTime time.Duration
	StatusCode  int
}

// SendHTTPRequest is used for sending a http request to given target
func SendHTTPRequest(url string, results chan Result) {
	start := time.Now()
	response, err := http.Get(url)

	if err != nil {
		log.Printf("Error calling: %v\n", url)
		return
	}

	httpResult := Result{URL: url, ElapsedTime: time.Since(start), StatusCode: response.StatusCode}

	results <- httpResult
}
