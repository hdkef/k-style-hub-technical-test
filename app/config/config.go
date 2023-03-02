package config

import (
	"errors"
	"os"
	"os/exec"
)

type AppConfig struct {
	DB_USERNAME string
	DB_PASS     string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	JWT_SECRET  string
	APP_PORT    string
}

func NewAPPConfig() *AppConfig {
	var config AppConfig
	config.loadENV()
	return &config
}

func (u *AppConfig) loadENV() {
	//source .env
	exec.Command("source", ".env")

	db_username, exist := os.LookupEnv("DB_USERNAME")
	if !exist {
		panic(errors.New("ENV DB_USERNAME MISSING"))
	}

	db_pass, exist := os.LookupEnv("DB_PASS")
	if !exist {
		panic(errors.New("ENV DB_PASS MISSING"))
	}

	db_host, exist := os.LookupEnv("DB_HOST")
	if !exist {
		panic(errors.New("ENV DB_HOST MISSING"))
	}

	db_port, exist := os.LookupEnv("DB_PORT")
	if !exist {
		panic(errors.New("ENV DB_PORT MISSING"))
	}

	db_name, exist := os.LookupEnv("DB_NAME")
	if !exist {
		panic(errors.New("ENV DB_NAME MISSING"))
	}

	jwt_secret, exist := os.LookupEnv("JWT_SECRET")
	if !exist {
		panic(errors.New("ENV JWT_SECRET MISSING"))
	}

	app_port, exist := os.LookupEnv("APP_PORT")
	if !exist {
		panic(errors.New("ENV APP_PORT MISSING"))
	}

	u.DB_HOST = db_host
	u.DB_NAME = db_name
	u.DB_USERNAME = db_username
	u.DB_PASS = db_pass
	u.DB_PORT = db_port
	u.JWT_SECRET = jwt_secret
	u.APP_PORT = app_port
}
