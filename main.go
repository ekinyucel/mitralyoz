package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/ekinyucel/mitralyoz/http"
	"github.com/ekinyucel/mitralyoz/work"
)

const (
	loadTestTime time.Duration = 30
	users                      = 10 // The total amount of concurrent users
	rampPeriod                 = 2  // Linear ramp-up period to create users
)

func createUsers(wg *sync.WaitGroup, results chan http.Result) {
	for i := 1; i <= users; i++ {
		wg.Add(1)

		go work.DoWork(i, wg, results)

		time.Sleep(time.Duration(int(1000*rampPeriod)) * time.Millisecond)
	}
	fmt.Println("all users are created")
}

func gatherResults(results chan http.Result) {
	for {
		select {
		case result := <-results:
			fmt.Printf("url %v statuscode %v total elapsed time %v\n", result.URL, result.StatusCode, result.ElapsedTime)
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	wg := sync.WaitGroup{}

	endTime := time.Now().Add(time.Second * loadTestTime)

	results := make(chan http.Result)

	go gatherResults(results)
	go createUsers(&wg, results)

	for range time.Tick(1 * time.Second) {
		if endTime.Before(time.Now()) {
			wg.Wait()
			fmt.Println("the load test has finished")
			break
		}
	}
}
