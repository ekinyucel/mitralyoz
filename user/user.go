package user

import (
	"log"
	"sync"
	"time"

	"github.com/ekinyucel/mitralyoz/config"
	"github.com/ekinyucel/mitralyoz/http"
	"github.com/ekinyucel/mitralyoz/work"
)

// CreateUsers function is used for generating user for the current test run.
// Function accepts test configuration details, waitgroup to be assigned to each user and results channel for sending results
func CreateUsers(testConfig config.TestConfig, wg *sync.WaitGroup, results chan http.Result) {
	users := testConfig.LoadTest.Users
	rampUpTime := testConfig.LoadTest.Rampup

	for i := 1; i <= users; i++ {
		wg.Add(1)

		go work.DoWork(testConfig, i, wg, results)

		time.Sleep(time.Duration(int(1000*rampUpTime)) * time.Millisecond)
	}
	log.Println("all users are created")
}
