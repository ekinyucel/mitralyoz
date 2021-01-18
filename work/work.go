package work

import (
	"fmt"
	"sync"

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
			fmt.Printf("User %d is started %d work\n", userID, usecase.ID)

			http.SendHTTPRequest(usecase.URL, results)

			fmt.Printf("User %d is finished %d work\n", userID, usecase.ID)
		}
	}
}
