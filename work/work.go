package work

import (
	"fmt"
	"sync"

	"github.com/ekinyucel/mitralyoz/config"
	"github.com/ekinyucel/mitralyoz/http"
	"github.com/ekinyucel/mitralyoz/usecase"
)

// Work represents a user action
type Work struct {
	workID  int
	useCase usecase.UseCase
}

// DoWork is responsible for executing requests for each user
func DoWork(testConfig config.TestConfig, userID int, wg *sync.WaitGroup, results chan http.Result) {
	defer wg.Done()
	useCases := usecase.InitializeUseCase()
	iterations := testConfig.LoadTest.Iterations

	for i := 0; i < iterations; i++ {
		for _, usecase := range useCases {
			fmt.Printf("User %d is started %d work\n", userID, usecase.UseCaseID)

			http.SendHTTPRequest(usecase.Url, results)

			fmt.Printf("User %d is finished %d work\n", userID, usecase.UseCaseID)
		}
	}
}
