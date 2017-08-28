package main

import (
	"POC/config"
	"POC/restclient"
	"log"

	"github.com/ContinuumLLC/platform-common-lib/src/kafka"
)

var restClient restclient.RestClient = restclient.GetClientFactoryInstance().GetRestClient(2)

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

func StartConsumer() {
	consumer, err := getKafkafkaConsumer()
	if err != nil {
		log.Printf("Error creating Consumer")
	}

	consumer.PullHandler(func(conMessage kafka.ConsumerMessage) {

		log.Println(conMessage.Message)
		log.Println(conMessage.Partition)
		log.Println(conMessage.Topic)
		log.Println(conMessage.Offset)

		queryParams := (map[string]string{
			"id": "John",
		})
		url := "http://localhost:8080/"
		body, err := restClient.Get(url, queryParams)
		if err != nil {
			log.Printf(err.Error())
		}

		log.Println(body)
	})

}
