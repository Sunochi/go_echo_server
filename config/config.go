package config

import (
	"encoding/json"
	"io/ioutil"
)

type config struct {
	DB  dbConfig  `json:"db"`
	AWS awsConfig `json:"aws"`
}

type dbConfig struct {
	Test     DBSetting `json:"test"`
	TestRead DBSetting `json:"test_read"`
}

type DBSetting struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
}

type awsConfig struct {
	Region string `json:"region"`
}

var c *config

func Init(filename *string) {

	jsonConfig, err := ioutil.ReadFile(*filename)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(jsonConfig, &c)
	if err != nil {
		panic(err)
	}
}

func GetConfig() *config {
	return c
}
