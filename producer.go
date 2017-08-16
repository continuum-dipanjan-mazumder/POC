package main

import (
	"POC/config"

	"github.com/ContinuumLLC/platform-common-lib/src/kafka"
)

func getKafkaProducer() (kafka.ProducerService, error) {

	producerconfig := new(kafka.ProducerConfig)

	producerconfig.BrokerAddress = []string{config.GetBrokerAddress()}
	producerFactory := new(kafka.ProducerFactoryImpl)
	producer, err := producerFactory.GetProducerService(*producerconfig)

	if err != nil {
		return nil, err
	}
	return producer, nil

}
