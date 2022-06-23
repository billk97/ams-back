package utils

import (
	"io/ioutil"
	"log"

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
	DB DB `yaml:"db"`
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
