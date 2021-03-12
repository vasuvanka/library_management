package config

import (
	env "github.com/Netflix/go-env"
)
type Config struct {
	Port string `env:"PORT,default=3000"`
	DbURI string `env:"DBURI,default=mongodb://localhost:27017/library"`
}

func NewConfig() *Config {
	return &Config{}
}

func (conf *Config) Init() error {
	_, err := env.UnmarshalFromEnviron(conf);
	return err
}