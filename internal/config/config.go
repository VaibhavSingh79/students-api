package config

import (
	"flag"
	"log"
	"os"

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

	configPath = os.Getenv("CONFIG_PATH") // get config path from env, os.Getenv is used to get the value of the environment variable.

	if configPath == "" {
		flags := flag.String("config", "", "path to the config file ") // flag.String is used to define a string flag with specified name, default value, and usage string.
		flag.Parse()                                                   // flag.Parse is used to parse the command-line flags.

		configPath = *flags // get config path from flag.

		if configPath == "" {
			log.Fatal("Config path is not set") // log.Fatal is used to log the error message and exit the program.
		}
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath) //fatalf is used to log the error message and exit the program.
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("cannot read config file: %s", err.Error())
	}
	return &cfg
}
