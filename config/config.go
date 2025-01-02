package config

import (
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2/log"
	"gopkg.in/yaml.v2"
)

type Config struct {
	APP struct {
		Address           string `yaml:"address"`
		PREFIX            string `yaml:"prefix"`
		NAME              string `yaml:"name"`
		ENABLEPRINTROUTES bool   `yaml:"enablePrintRoutes"`
		SERVERHEADER      string `yaml:"SeverHeader"`
		AllowOrigins      string `yaml:"allowOrigins"`
		AllowMethods      string `yaml:"allowMethods"`
		AllowHeaders      string `yaml:"allowHeaders"`
		AllowCredentials  bool   `yaml:"allowCredentials"`
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

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		log.Errorf("Didnt decode it right: %v", err)
		log.Infof("the decoded err: %v", cfg)
		return nil, err
	}

	/* if err := yaml.Unmarshal(file, cfg); err != nil {
		log.Fatalf("Unmarshal: %v", err)
	} */

	log.Infof("Config yml has been decoded: %v", cfg)

	return cfg, nil

}
