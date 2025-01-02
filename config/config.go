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
		PREFIX            string `ymal:"prefix"`
		NAME              string `ymal:"name"`
		ENABLEPRINTROUTES bool   `ymal:"enablePrintRoutes"`
		SERVERHEADER      string `ymal:"SeverHeader"`
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

	file, err := os.ReadFile(fp)
	if err != nil {
		return nil, err
	}
	//defer file.Close()

	cfg := &Config{}

	/* decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		return nil, err
	} */

	if err := yaml.Unmarshal(file, cfg); err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	log.Infof("Config yml has been decoded: %v", cfg)

	return cfg, nil

}
