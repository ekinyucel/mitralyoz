package usecase

// UseCase struct is for holding a single test case
type UseCase struct {
	UseCaseID int
	Url       string
}

// InitializeUseCase is used for creating use cases for performance test
// This function returns a slice of uses cases
func InitializeUseCase() []UseCase {
	useCase := UseCase{UseCaseID: 1, Url: "https://google.com"}

	return []UseCase{useCase}
}
