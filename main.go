package main

import (
	"fmt"
	"sync"
	"time"
)

func work(userID int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("User %d is created\n", userID)

	fmt.Printf("User %d has finished\n", userID)
}

func main() {
	users := 10
	rampPeriod := 1 // seconds

	wg := sync.WaitGroup{}

	startTime := time.Now()
	for i := 1; i <= users; i++ {
		wg.Add(1)

		go work(i, &wg)

		time.Sleep(time.Duration(int(1000*rampPeriod)) * time.Millisecond)
	}
	endTime := time.Now()

	fmt.Printf("Load test started on %s \n", startTime)
	fmt.Printf("Load test ended on %s \n", endTime)

	wg.Wait()
}
