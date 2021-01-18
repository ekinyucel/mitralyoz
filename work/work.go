package work

import (
	"fmt"
	"sync"

	"github.com/ekinyucel/mitralyoz/http"
	"github.com/ekinyucel/mitralyoz/usecase"
)

const (
	iterations = 5 // The iteration amount for each user to execute the actions
)

// Work represents a user action
type Work struct {
	workID  int
	useCase usecase.UseCase
}

// DoWork is responsible for executing requests for each user
func DoWork(userID int, wg *sync.WaitGroup, results chan http.Result) {
	defer wg.Done()
	useCases := usecase.InitializeUseCase()

	for i := 0; i < iterations; i++ {
		for _, usecase := range useCases {
			fmt.Printf("User %d is started %d work\n", userID, usecase.UseCaseID)

			http.SendHTTPRequest(usecase.Url, results)

			fmt.Printf("User %d is finished %d work\n", userID, usecase.UseCaseID)
		}
	}
}
