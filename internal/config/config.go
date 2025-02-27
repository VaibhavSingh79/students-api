package config

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address" env-required:"true"`
}

// env-default:"production"
type Config struct {
	ENV         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string

	// 1. Get CONFIG_PATH from environment
	configPath = os.Getenv("CONFIG_PATH")

	// 2. Get --config flag if CONFIG_PATH is not set
	if configPath == "" {
		flags := flag.String("config", "", "path to the config file")
		flag.Parse()
		configPath = *flags
	}

	// 3. Default fallback to root-based relative path (config/local.yaml)
	if configPath == "" {
		// Find the project's root directory (assumes binary is in cmd/stu-api/)
		exePath, err := os.Executable()
		if err != nil {
			log.Fatalf("Failed to determine executable path: %v", err)
		}
		projectRoot := filepath.Join(filepath.Dir(exePath), "../../")
		configPath = filepath.Join(projectRoot, "config/local.yaml")
	}

	// Check if the config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", configPath)
	}

	// Parse the config with cleanenv
	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("Cannot read config file: %s", err.Error())
	}

	return &cfg
}
