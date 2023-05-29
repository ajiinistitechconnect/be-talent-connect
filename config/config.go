package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

type ApiConfig struct {
	ApiHost string
	ApiPort string
}

type Config struct {
	DbConfig
	ApiConfig
}

func (c *Config) ReadConfigFile() error {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println(err)
		return errors.New("Failed to load .env file")
	}

	c.DbConfig = DbConfig{
		Host:     os.Getenv("DBHOST"),
		Port:     os.Getenv("DBPORT"),
		Name:     os.Getenv("DBNAME"),
		User:     os.Getenv("DBUSER"),
		Password: os.Getenv("DBPASS"),
	}

	c.ApiConfig = ApiConfig{
		ApiHost: os.Getenv("API_HOST"),
		ApiPort: os.Getenv("API_PORT"),
	}

	if c.DbConfig.Host == "" || c.DbConfig.Name == "" || c.DbConfig.Password == "" || c.DbConfig.Port == "" || c.DbConfig.User == "" || c.ApiConfig.ApiHost == "" || c.ApiConfig.ApiPort == "" {
		return errors.New("Missing required field")
	}
	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cfg.ReadConfigFile()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
