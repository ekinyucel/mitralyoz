package work

import (
	"log"
	"sync"
	"time"

	"github.com/ekinyucel/mitralyoz/config"
	"github.com/ekinyucel/mitralyoz/http"
	"github.com/ekinyucel/mitralyoz/usecase"
)

// DoWork is responsible for executing requests for each user
func DoWork(testConfig config.TestConfig, userID int, wg *sync.WaitGroup, results chan http.Result) {
	defer wg.Done()
	useCases := usecase.InitializeUseCase()
	iterations := testConfig.LoadTest.Iterations

	for i := 0; i < iterations; i++ {
		for _, usecase := range useCases.UseCase {
			log.Printf("User %d is started, use case %d\n", userID, usecase.ID)

			http.SendHTTPRequest(usecase.URL, results)

			log.Printf("User %d is finished, use case %d\n", userID, usecase.ID)

			time.Sleep(time.Duration(int(1000*testConfig.LoadTest.Cooldown)) * time.Millisecond)
		}
	}
}
