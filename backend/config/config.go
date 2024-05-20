package config

import (
	"log"

	"github.com/caarlos0/env/v10"
)

type Config struct {
	ClientDomain    string `env:"CLIENT_DOMAIN" envDefault:"localhost:3000"`
	ClientProto     string `env:"CLIENT_PROTO"  envDefault:"http://"`
	DatabaseUser    string `env:"DATABASE_USER" envDefault:"root"`
	DatabasePass    string `env:"DATABASE_PASS" envDefault:"mysql"`
	DatabaseHost    string `env:"DATABASE_HOST" envDefault:"localhost"`
	DatabasePort    int    `env:"DATABASE_PORT" envDefault:"3306"`
	DatabaseName    string `env:"DATABASE_NAME" envDefault:"twitter_clone_app_development"`
	DatabaseSslCert string `env:"DATABASE_SSL_CERT"`
}

func New() Config {
	var config Config
	if error := env.Parse(&config); error != nil {
		log.Fatal(error)
		panic(error)
	}
	return config
}

func (c *Config) ClientHost() string {
	return c.ClientProto + c.ClientDomain
}
