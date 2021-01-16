package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	loadTestTime time.Duration = 30
	workCount                  = 100
)

// Work represents a user action
type Work struct {
	workID int
}

func doWork(userID int, works <-chan Work, wg *sync.WaitGroup) {
	defer wg.Done()
	for w := range works {
		fmt.Printf("User %d is started %d work\n", userID, w.workID)
		time.Sleep(time.Millisecond * 1800)
		fmt.Printf("User %d is finished %d work\n", userID, w.workID)
	}
	fmt.Println("User ", userID, " is done")
}

func createUsers(works <-chan Work, wg *sync.WaitGroup) {
	users := 10
	rampPeriod := 5 // seconds

	for i := 1; i <= users; i++ {
		wg.Add(1)

		go doWork(i, works, wg)

		time.Sleep(time.Duration(int(1000*rampPeriod)) * time.Millisecond)
	}
	fmt.Println("all users are created")
}

func createWorkLoad(works chan<- Work, workCount int) {
	for i := 1; i < workCount; i++ {
		works <- Work{workID: i}
	}
	fmt.Println("all works are created")
}

func main() {
	wg := sync.WaitGroup{}

	endTime := time.Now().Add(time.Second * loadTestTime)

	works := make(chan Work, workCount)
	go createWorkLoad(works, workCount)

	for range time.Tick(1 * time.Second) {
		if endTime.Before(time.Now()) {
			close(works)
			wg.Wait()
			fmt.Println("the load test has finished")
			break
		}

		go createUsers(works, &wg)
	}
}
