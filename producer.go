package main

import (
	"github.com/ContinuumLLC/platform-common-lib/src/kafka"
)

func getKafkaProducer() (kafka.ProducerService, error) {
	producerconfig := new(kafka.ProducerConfig)
	producerconfig.BrokerAddress = []string{"localhost:9092"}
	producerFactory := new(kafka.ProducerFactoryImpl)
	producer, err := producerFactory.GetProducerService(*producerconfig)

	if err != nil {
		return nil, err
	}
	return producer, nil

}
