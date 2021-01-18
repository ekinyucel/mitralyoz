package usecase

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Config struct holds relevant details about each use case
type Config struct {
	UseCase []struct {
		ID  int    `yaml:"id"`
		URL string `yaml:"url"`
	} `yaml:"usecase"`
}

// InitializeUseCase is used for creating use cases for performance test
// This function returns a slice of uses cases
func InitializeUseCase() *Config {
	f, err := os.Open("resources/usecase.yml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	config := &Config{}
	decoder := yaml.NewDecoder(f)

	if err := decoder.Decode(&config); err != nil {
		panic(err)
	}

	return config
}
