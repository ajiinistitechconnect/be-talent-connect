package config

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
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

type SMTPConfig struct {
	SMTPHost       string
	SMTPPort       string
	SMTPSenderName string
	SMTPEmail      string
	SMTPPassword   string
}

type TokenConfig struct {
	ApplicationName     string
	JwtSignatureKey     string
	JwtSigningMethod    *jwt.SigningMethodHMAC
	AccessTokenLifeTime time.Duration
}

type RedisConfig struct {
	Address  string
	Password string
	Db       int
}

type Config struct {
	DbConfig
	ApiConfig
	SMTPConfig
	TokenConfig
	RedisConfig
}

func (c *Config) ReadConfigFile() error {
	// if os.Getenv("ENV") == "local" {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
		return errors.New("Failed to load .env file")
	}
	// }

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

	if os.Getenv("PORT") != "" {
		c.ApiConfig.ApiPort = os.Getenv("PORT")
	}

	// c.SMTPConfig = SMTPConfig{
	// 	SMTPHost:       os.Getenv("SMTP_HOST"),
	// 	SMTPPort:       os.Getenv("SMTP_PORT"),
	// 	SMTPSenderName: os.Getenv("SMTP_SENDER"),
	// 	SMTPEmail:      os.Getenv("SMTP_EMAIL"),
	// 	SMTPPassword:   os.Getenv("SMTP_PASS"),
	// }

	c.TokenConfig = TokenConfig{
		ApplicationName:     "TALENTCONNECT",
		JwtSignatureKey:     "x/A?D(G+KaPdSgVkYp3s6v9y$B&E)H@M",
		JwtSigningMethod:    jwt.SigningMethodHS256,
		AccessTokenLifeTime: time.Minute * 5,
	}

	c.RedisConfig = RedisConfig{
		Address:  os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		Db:       0,
	}
	// c.SMTPEmail == "" || c.SMTPHost == "" || c.SMTPPassword == "" || c.SMTPPort == "" || c.SMTPSenderName == ""

	if c.DbConfig.Host == "" || c.DbConfig.Name == "" || c.DbConfig.Password == "" || c.DbConfig.Port == "" || c.DbConfig.User == "" {
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
