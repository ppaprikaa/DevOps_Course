package cfg

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
)

func Load() *Cfg {
	var (
		Cfg     Cfg
		CfgPath = os.Getenv("CONFIG_PATH")
	)

	if strings.TrimSpace(CfgPath) == "" {
		CfgPath = "cfg/cfg.yaml"
	}

	if err := cleanenv.ReadConfig(CfgPath, &Cfg); err != nil && !errors.Is(err, os.ErrNotExist) {
		log.Fatal(err)
	}

	if err := cleanenv.ReadEnv(&Cfg); err != nil {
		log.Fatal(err)
	}

	return &Cfg
}

type Cfg struct {
	Server Server `json:"server" yaml:"server" toml:"server"`
	DB     DB     `json:"db" yaml:"db" toml:"db"`
}

type Server struct {
	Host string `json:"host" yaml:"host" toml:"host" env:"SERVER_HOST" env-default:""`
	Port int    `json:"port" yaml:"port" toml:"port" env:"SERVER_PORT" env-default:"5000"`
}

type DB struct {
	DSN string `json:"dsn" yaml:"dsn" toml:"dsn" env:"DB_DSN"`
}
