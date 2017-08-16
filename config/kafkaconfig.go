package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Kafkaconfig struct {
	TopicName string `json:"topic"`
}

func LoadConfiguration(file string) Kafkaconfig {
	var config Kafkaconfig
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
