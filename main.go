package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/ekinyucel/mitralyoz/http"
	"github.com/ekinyucel/mitralyoz/usecase"
)

const (
	loadTestTime time.Duration = 30
	users                      = 10 // The total amount of concurrent users
	rampPeriod                 = 2  // Linear ramp-up period to create users
	iterations                 = 5  // The iteration amount for each user to execute the actions
)

// Work represents a user action
type Work struct {
	workID  int
	useCase usecase.UseCase
}

func doWork(userID int, wg *sync.WaitGroup) {
	defer wg.Done()
	useCases := usecase.InitializeUseCase()

	for i := 0; i < iterations; i++ {
		for _, usecase := range useCases {
			fmt.Printf("User %d is started %d work\n", userID, usecase.UseCaseID)

			http.SendHTTPRequest(usecase.Url)

			fmt.Printf("User %d is finished %d work\n", userID, usecase.UseCaseID)
		}
	}
}

func createUsers(wg *sync.WaitGroup) {
	for i := 1; i <= users; i++ {
		wg.Add(1)

		go doWork(i, wg)

		time.Sleep(time.Duration(int(1000*rampPeriod)) * time.Millisecond)
	}
	fmt.Println("all users are created")
}

func main() {
	wg := sync.WaitGroup{}

	endTime := time.Now().Add(time.Second * loadTestTime)

	go createUsers(&wg)

	for range time.Tick(1 * time.Second) {
		if endTime.Before(time.Now()) {
			wg.Wait()
			fmt.Println("the load test has finished")
			break
		}
	}
}
