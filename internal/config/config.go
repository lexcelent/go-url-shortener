package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env  string `yaml:"env" env-default:"dev"`
	Host string `yaml:"host" env-default:"localhost"`
	Port int    `yaml:"port" env-default:"8080"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("Не задан путь до конфигурации")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("Файла конфига не существует: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "Путь до конфигурационного файла")

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
