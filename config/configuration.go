package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// TestConfig holds configuration values
type TestConfig struct {
	LoadTest struct {
		Users      int `yaml:"users"`
		Cooldown   int `yaml:"cooldown"`
		Rampup     int `yaml:"rampup"`
		TotalTime  int `yaml:"totalTime"`
		Iterations int `yaml:"iterations"`
	} `yaml:"loadtest"`
}

// ReadTestConfig function reads configuration values from yaml
func ReadTestConfig() *TestConfig {
	f, err := os.Open("resources/config.yml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	testConfig := &TestConfig{}
	decoder := yaml.NewDecoder(f)

	if err := decoder.Decode(&testConfig); err != nil {
		panic(err)
	}

	return testConfig
}
