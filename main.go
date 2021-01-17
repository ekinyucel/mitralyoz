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
	users                      = 30
	rampPeriod                 = 5 // seconds
)

// Work represents a user action
type Work struct {
	workID  int
	useCase usecase.UseCase
}

func doWork(userID int, works <-chan Work, wg *sync.WaitGroup) {
	defer wg.Done()
	for w := range works {
		fmt.Printf("User %d is started %d work\n", userID, w.workID)

		http.SendHTTPRequest(w.useCase.Url)

		fmt.Printf("User %d is finished %d work\n", userID, w.workID)
	}
}

func createUsers(works <-chan Work, wg *sync.WaitGroup) {
	for i := 1; i <= users; i++ {
		wg.Add(1)

		go doWork(i, works, wg)

		time.Sleep(time.Duration(int(1000*rampPeriod)) * time.Millisecond)
	}
	fmt.Println("all users are created")
}

func main() {
	wg := sync.WaitGroup{}

	endTime := time.Now().Add(time.Second * loadTestTime)

	works := make(chan Work)

	go createUsers(works, &wg)

	for range time.Tick(1 * time.Second) {
		if endTime.Before(time.Now()) {
			close(works)
			wg.Wait()
			fmt.Println("the load test has finished")
			break
		}
		go func(works chan<- Work) {
			useCases := usecase.InitializeUseCase()
			works <- Work{workID: time.Now().Nanosecond(), useCase: useCases[0]}
		}(works)
	}
}
