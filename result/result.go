package result

import (
	"fmt"
	"time"

	"github.com/ekinyucel/mitralyoz/http"
)

// GatherResults function is used for listening the results of each use case run
// It accepts results channel
func GatherResults(results chan http.Result) {
	for {
		select {
		case result := <-results:
			fmt.Printf("url %v statuscode %v total elapsed time %v\n", result.URL, result.StatusCode, result.ElapsedTime)
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}
