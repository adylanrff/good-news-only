package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Config is the config object that is formed for the application configuration
type Config struct {
	Scrapper struct {
		Paralellism int `yaml:"parallelism"`
		SecondDelay int `yaml:"second_delay"`
	}
	Twitter struct {
		APIKey string `yaml:"apikey"`
	} `yaml:"twitter"`
	Reddit struct {
		Domain     string   `yaml:"domain"`
		URL        string   `yaml:"url"`
		Subreddits []string `yaml:"subreddits"`
		MaxPages   int      `yaml:"max_pages"`
	} `yaml:"reddit"`
	Postgre struct {
		URI string `yaml:"uri"`
	} `yaml:"postgre"`
}

// LoadConfig will load config file from the input configPath
func LoadConfig(configPath string) (*Config, error) {

	f, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
