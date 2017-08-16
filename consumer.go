package main

import (
	"POC/config"
	"fmt"

	"github.com/ContinuumLLC/platform-common-lib/src/kafka"
)

func getKafkafkaConsumer() (kafka.ConsumerService, error) {
	consumerconfig := new(kafka.ConsumerConfig)
	consumerconfig.BrokerAddress = []string{"localhost:9092"}
	conf := config.LoadConfiguration("config/kafka.json")
	fmt.Println()
	consumerconfig.Topics = []string{conf.TopicName}
	consumerconfig.GroupID = "dipanjan01"
	consumerFactory := new(kafka.ConsumerFactoryImpl)
	consumer, err := consumerFactory.GetConsumerService(*consumerconfig)

	if err != nil {
		return nil, err
	}
	return consumer, nil

}
