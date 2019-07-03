package config

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

type Config struct {
	DbHost     string
	DbPort     int
	DbName     string
	DbUsername string
	DbPassword string
}

func GetConfig() Config {
	config := Config{}
	err := gonfig.GetConf("config/config.json", &config)
	if err != nil {
		fmt.Println("DB Config file Error")
	}
	return config
}
