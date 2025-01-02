package config

import (
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2/log"
	"gopkg.in/yaml.v2"
)

type Config struct {
	APP struct {
		Address string `yaml:"address"`
		PREFIX  string `ymal:"prefix"`
	} `yaml:"app"`
	API struct {
		JsonPlaceholder string `yaml:"jsonplaceholder"`
	} `ymal:"api"`
}

func getProjectRoot() string {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// If we're in a subdirectory (like /config), we need to go up one level
	if filepath.Base(pwd) == "config" {
		return filepath.Dir(pwd)
	}

	return pwd
}

func LoadConfig() (*Config, error) {
	projectRoot := getProjectRoot()
	fp := filepath.Join(projectRoot, "config.yml")

	file, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	cfg := &Config{}

	log.Infof("Empty config: %v", cfg)

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		return nil, err
	}

	log.Infof("Config yml has been decoded: %v", cfg)

	return cfg, nil

}
