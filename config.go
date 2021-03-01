package authservice

import (
	"fmt"
	"os"
	"strconv"
)

// Config here goes the configuration variables for this app
type Config struct {
	DB struct {
		User     string
		DBName   string
		Password string
		Host     string
		Port     string
	}
	App struct {
		JWTKey     string
		JWTExpTime int
		Domain     string
	}
	Mail struct {
		Host     string
		Port     string
		From     string
		Password string
	}
}

// PostgresConnStr returns the postgres connection string
func (conf Config) PostgresConnStr() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		conf.DB.User, conf.DB.Password, conf.DB.Host, conf.DB.Port, conf.DB.DBName,
	)
}

// GetConfig reads the env variables and return them
func GetConfig() (Config, error) {
	var conf Config
	envKeys := []string{
		"DB_USER",
		"DB_NAME",
		"DB_PASSWORD",
		"DB_HOST",
		"DB_PORT",
		"JWT_KEY",
		"JWT_EXP_TIME",
		"APP_DOMAIN",
		"EMAIL_HOST",
		"EMAIL_PORT",
		"EMAIL_FROM",
		"EMAIL_PASSWORD",
	}
	for _, key := range envKeys {
		if err := checkEnv(key); err != nil {
			return conf, err
		}
	}
	conf.DB.User = os.Getenv("DB_USER")
	conf.DB.DBName = os.Getenv("DB_NAME")
	conf.DB.Password = os.Getenv("DB_PASSWORD")
	conf.DB.Host = os.Getenv("DB_HOST")
	conf.DB.Port = os.Getenv("DB_PORT")
	conf.App.JWTKey = os.Getenv("JWT_KEY")
	duration, err := strconv.Atoi(os.Getenv("JWT_EXP_TIME"))
	if err != nil {
		return conf, fmt.Errorf("JWT_EXP_TIME env variable must be a integer value: %s", err.Error())
	}
	conf.App.JWTExpTime = duration
	conf.App.Domain = os.Getenv("APP_DOMAIN")
	conf.Mail.Host = os.Getenv("EMAIL_HOST")
	conf.Mail.Port = os.Getenv("EMAIL_PORT")
	conf.Mail.From = os.Getenv("EMAIL_FROM")
	conf.Mail.Password = os.Getenv("EMAIL_PASSWORD")
	return conf, nil
}

func checkEnv(key string) error {
	value := os.Getenv(key)
	if value == "" {
		return fmt.Errorf("env variable %s is empty", key)
	}
	return nil
}
