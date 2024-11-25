package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type ServerConf struct {
	Port		string `envconfig:"SERVER_PORT" required:"true"`
	Endpoint	string `envconfig:"FAKE_ENDPOINT" required:"true"`
}

type DatabaseConf struct {
	Name	string `envconfig:"DB_NAME" required:"true"`
	User	string `envconfig:"DB_USER" required:"true"`
	Pass	string `envconfig:"DB_PASS" required:"true"`
	Adress	string `envconfig:"DB_ADRESS" required:"true"`
	Port	string `envconfig:"DB_PORT" required:"true"`
}

type Config struct {
	Serv	*ServerConf
	Db		*DatabaseConf
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	cfg := Config{}
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
