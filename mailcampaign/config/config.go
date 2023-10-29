package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App  App  `yaml:"app"`
	SMTP SMTP `yaml:"smtp"`
}

type App struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}

type SMTP struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	SenderName   string `yaml:"sender_name"`
	AuthEmail    string `yaml:"auth_email"`
	AuthPassword string `yaml:"auth_password"`
}

var Cfg *Config

func LoadConfig(filename string) (err error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	cfg := Config{}

	err = yaml.Unmarshal(f, &cfg)
	if err != nil {
		return
	}

	Cfg = &cfg
	return
}
