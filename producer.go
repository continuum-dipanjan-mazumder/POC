package main

import (
	"POC/config"
	"POC/input"
	"log"
	"time"

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

func StartProducer() {
	producer, err := getKafkaProducer()
	if err != nil {
		log.Printf("Error creating a Producer")
	}

	go pushData(producer, config.GetTopic())

}

func pushData(producer kafka.ProducerService, topic string) {

	i := 0
	for {
		time.Sleep(time.Second * 2)
		i = i + 1
		producer.Push(topic, input.GetMessage(string(i)))
	}
}
