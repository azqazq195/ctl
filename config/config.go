package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Name     string          `yaml:"name"`
	Services ServiceCategory `yaml:"services"`
}

type ServiceCategory struct {
	Downloads map[string]DownloadService `yaml:"downloads"`
	Installs  map[string]InstallService  `yaml:"installs"`
	Runs      map[string]RunService      `yaml:"runs"`
}

type DownloadService struct {
	Description string   `yaml:"description"`
	URLs        []string `yaml:"urls"`
	Required    bool     `yaml:"required"`
}

type InstallService struct {
	Path string `yaml:"path"`
}

type RunService struct {
	Path string `yaml:"path"`
}

func LoadConfig() Config {
	configFile, err := os.ReadFile("/Users/seongha.moon/Documents/development/project/ctl/config.yml")
	//configFile, err := os.ReadFile("config.yml")
	if err != nil {
		fmt.Println("Error reading config file:", err)
		os.Exit(1)
	}

	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Println("Error parsing config file:", err)
		os.Exit(1)
	}

	return config
}
