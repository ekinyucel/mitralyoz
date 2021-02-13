package result

import (
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ekinyucel/mitralyoz/http"
)

type syncMap struct {
	mutex sync.Mutex
	m     map[int]int32
}

func newSyncMap() *syncMap {
	return &syncMap{
		m: make(map[int]int32),
	}
}

func (s *syncMap) Set(key int, value int32) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.m[key] = value
}

func (s *syncMap) Get(key int) int32 {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.m[key]
}

var totalRequests *syncMap
var elapsedTimeSum *int32 = new(int32)

// GatherResults function is used for listening the results of each use case run
// It accepts results channel
func GatherResults(results chan http.Result) {
	totalRequests = newSyncMap()
	for {
		select {
		case result := <-results:
			requestCount := totalRequests.Get(result.UseCaseID) + 1
			totalRequests.Set(result.UseCaseID, requestCount)
			atomic.AddInt32(elapsedTimeSum, int32(result.ElapsedTime))

			log.Printf("Request %v : url %v statuscode %v total elapsed time %v\n",
				requestCount, result.URL, result.StatusCode, result.ElapsedTime)
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// GetTotalRequestCount function returns the total count of requests have been executed as part of the load test run
func GetTotalRequestCount() {
	for key := range totalRequests.m {
		log.Printf("total request count of use case %v %v", key, totalRequests.Get(key))
	}
}
