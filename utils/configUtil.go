package utils

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type DB struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

type Env struct {
	DB    DB `yaml:"db"`
	Aries string
}

var Config Env

func InitYamlConfig() {
	file, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatal("Could not find config.yaml")
		return
	}
	error := yaml.Unmarshal([]byte(file), &Config)
	if error != nil {
		log.Fatal("Failed to parce: Invaled yaml syntax!")
		return
	}
}

func InitEnv() {
	godotenv.Load(".env")
	Config.DB.Name = os.Getenv("DB_NAME")
	Config.DB.Port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	Config.DB.Host = os.Getenv("DB_HOST")
	Config.DB.Username = os.Getenv("DB_USERNAME")
	Config.DB.Password = os.Getenv("DB_PASSWORD")
	Config.Aries = os.Getenv("ARIES_HOST")
}