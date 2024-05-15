package config

import (
	"flag"
	"github.com/lpernett/godotenv"
	"log"
	"os"
)

type Config struct {
	Server   Server
	Database Database
	Secret   Secret
}

type Server struct {
	Host string
	Port string
}

type Database struct {
	Host string
	Port string
	User string
	Pass string
	Name string
	Tz   string
}

type Secret struct {
	Jwt string
}

func Get() *Config {
	fileFlag := flag.String("env", "", "file .env location path absolute")
	flag.Parse()

	var err error
	if *fileFlag != "" {
		err = godotenv.Load(*fileFlag)
	} else {
		err = godotenv.Load()
	}

	if err != nil {
		log.Fatal("error when load .env: ", err.Error())
	}

	return &Config{
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Database: Database{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Name: os.Getenv("DB_NAME"),
			Tz:   os.Getenv("DB_TZ"),
		},
		Secret: Secret{
			Jwt: os.Getenv("SECRET_JWT"),
		},
	}
}
