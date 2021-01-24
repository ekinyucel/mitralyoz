package main

import (
	"log"
	"sync"
	"time"

	"github.com/ekinyucel/mitralyoz/config"
	"github.com/ekinyucel/mitralyoz/http"
	"github.com/ekinyucel/mitralyoz/result"
	"github.com/ekinyucel/mitralyoz/user"
)

func main() {
	testConfig := config.ReadTestConfig()

	wg := sync.WaitGroup{}

	endTime := time.Now().Add(time.Second * time.Duration(testConfig.LoadTest.TotalTime))

	results := make(chan http.Result)

	go result.GatherResults(results)

	wg.Add(testConfig.LoadTest.Users)
	go user.CreateUsers(*testConfig, &wg, results)

	for range time.Tick(1 * time.Second) {
		if endTime.Before(time.Now()) {
			wg.Wait()
			log.Printf("total request count %v", result.GetTotalRequestCount())
			log.Println("the load test has finished")
			break
		}
	}
}
