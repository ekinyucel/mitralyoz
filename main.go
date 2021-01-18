package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/ekinyucel/mitralyoz/config"
	"github.com/ekinyucel/mitralyoz/http"
	"github.com/ekinyucel/mitralyoz/work"
)

func createUsers(testConfig config.TestConfig, wg *sync.WaitGroup, results chan http.Result) {
	users := testConfig.LoadTest.Users
	rampUpTime := testConfig.LoadTest.Rampup

	for i := 1; i <= users; i++ {
		wg.Add(1)

		go work.DoWork(testConfig, i, wg, results)

		time.Sleep(time.Duration(int(1000*rampUpTime)) * time.Millisecond)
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
	testConfig := config.ReadConfig()

	wg := sync.WaitGroup{}

	endTime := time.Now().Add(time.Second * time.Duration(testConfig.LoadTest.TotalTime))

	results := make(chan http.Result)

	go gatherResults(results)
	go createUsers(*testConfig, &wg, results)

	for range time.Tick(1 * time.Second) {
		if endTime.Before(time.Now()) {
			wg.Wait()
			fmt.Println("the load test has finished")
			break
		}
	}
}
