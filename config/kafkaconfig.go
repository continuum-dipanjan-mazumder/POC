package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Kafkaconfig struct {
	Hostname  string `json:"hostname"`
	Port      int    `json:"port"`
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

func GetBrokerAddress() string {
	conf := LoadConfiguration("config/kafka.json")
	address := fmt.Sprintf("%s%s%d", conf.Hostname, ":", conf.Port)
	return address
}

func GetTopic() string {
	conf := LoadConfiguration("config/kafka.json")
	return conf.TopicName
}
