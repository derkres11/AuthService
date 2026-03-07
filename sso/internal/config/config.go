package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string     `yaml:"env" env-default:"local"`
	StoragePath string     `yaml:"storage_path" env-required:"true"`
	GRPC        GRPCConfig `yaml:"grpc"`
}

type GRPCConfig struct {
	Port     int           `yaml:"port" env-default:"44044"`
	Timeout  string        `yaml:"timeout" env-default:"5s"`
	TokenTTL time.Duration `yaml:"token_ttl" env-default:"1h"`
	Host     string        `yaml:"host" env-default:"localhost"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is not set")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist at path: " + path)
	}

	var cfg Config

	cleanenv.ReadConfig(path, &cfg)

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "config/local.yaml", "Path to the config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}
