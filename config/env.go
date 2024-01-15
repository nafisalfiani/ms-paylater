package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Value struct {
	Database Database
	Auth     Auth
	Log      Log
	Server   Server
}

type Database struct {
	DbUrl      string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
}

type Auth struct {
	SecretKey string
}

type Server struct {
	Base string
	Port int
}

type Log struct {
	Level string
}

func InitEnv() (*Value, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		return nil, err
	}

	return &Value{
		Database: Database{
			DbUrl:      os.Getenv("DB_URL"),
			DbPort:     os.Getenv("DB_PORT"),
			DbName:     os.Getenv("DB_NAME"),
			DbUser:     os.Getenv("DB_USER"),
			DbPassword: os.Getenv("DB_PASSWORD"),
		},
		Auth: Auth{
			SecretKey: os.Getenv("AUTH_SECRETKEY"),
		},
		Log: Log{
			Level: os.Getenv("LOG_LEVEL"),
		},
		Server: Server{
			Base: os.Getenv("SERVER_BASE"),
			Port: port,
		},
	}, nil
}
