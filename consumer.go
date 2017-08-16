package main

import (
	"POC/config"

	"github.com/ContinuumLLC/platform-common-lib/src/kafka"
)

func getKafkafkaConsumer() (kafka.ConsumerService, error) {

	consumerconfig := new(kafka.ConsumerConfig)
	consumerconfig.BrokerAddress = []string{config.GetBrokerAddress()}
	consumerconfig.Topics = []string{config.GetTopic()}
	consumerconfig.GroupID = "dipanjan01"
	consumerFactory := new(kafka.ConsumerFactoryImpl)
	consumer, err := consumerFactory.GetConsumerService(*consumerconfig)

	if err != nil {
		return nil, err
	}
	return consumer, nil

}
