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
	conf := LoadConfiguration(getConfigFileLocation() + "/kafka.json")

	address := fmt.Sprintf("%s%s%d", conf.Hostname, ":", conf.Port)
	fmt.Println("Address is %s", address)
	return address
}

func GetTopic() string {
	conf := LoadConfiguration(getConfigFileLocation() + "/kafka.json")
	return conf.TopicName
}

func getConfigFileLocation() string {
	return os.Getenv("KAFKA_CONFIG")
}
