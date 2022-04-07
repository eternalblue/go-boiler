package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"io/ioutil"
	"os"
)

type Configuration struct {
	LogLevel      string `yaml:"log_level"`
	EggRepository struct {
		Memory struct {
			Size int `yaml:"size"`
		} `yaml:"memory"`
	} `yaml:"egg_repository"`
}

// InitConfiguration load configuration from the correct yaml.
func InitConfiguration() *Configuration {
	cfgPath := fmt.Sprintf("%s/application-%s.yaml", getConfigFolderPath(), getEnv())

	logrus.Info("loding configuration", cfgPath)

	cfg, err := load(cfgPath)
	if err != nil {
		panic(err)
	}

	return cfg
}

func IsProduction() bool {
	return os.Getenv("GO_ENVIRONMENT") == "production"
}

func load(path string) (*Configuration, error) {
	var configuration Configuration

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	yamlFile = []byte(os.ExpandEnv(string(yamlFile)))

	err = yaml.Unmarshal(yamlFile, &configuration)
	if err != nil {
		return nil, err
	}

	return &configuration, nil
}

func getConfigFolderPath() string {
	pwd, _ := os.Getwd()
	return pwd + "/config"
}

func getEnv() string {
	environment := os.Getenv("GO_ENVIRONMENT")
	if environment == "" {
		environment = "local"
	}

	return environment
}
