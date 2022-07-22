package utils

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DB struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

type EmailVars struct {
	Host     string
	From     string
	Password string
	Username string
}

type Env struct {
	DB         DB `yaml:"db"`
	Aries      string
	EmailVars  EmailVars
	JwtSecret  string
	HostDomain string
}

var Config Env

func InitEnv() {
	godotenv.Load(".env")
	Config.DB.Name = os.Getenv("DB_NAME")
	Config.DB.Port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	Config.DB.Host = os.Getenv("DB_HOST")
	Config.DB.Username = os.Getenv("DB_USERNAME")
	Config.DB.Password = os.Getenv("DB_PASSWORD")
	Config.Aries = os.Getenv("ARIES_HOST")
	Config.EmailVars = EmailVars{
		Host:     os.Getenv("EMAIL_HOST"),
		From:     os.Getenv("EMAIL_FROM"),
		Password: os.Getenv("EMAIL_PASSWORD"),
		Username: os.Getenv("EMAIL_USERNAME"),
	}
	Config.JwtSecret = os.Getenv("JWT_SECRET")
	Config.HostDomain = os.Getenv("HOST_DOMAIN")
	// todo add checks if empty
}
