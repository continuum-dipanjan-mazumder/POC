package main

import (
	"encoding/json"
	"fmt"
	"kafkaPOC/config"
	"log"
	"time"

	"github.com/ContinuumLLC/platform-common-lib/src/kafka"
)

var cha1 = make(chan string)

func main() {

	producer, err := getKafkaProducer()
	if err != nil {
		log.Printf("Error creating a Producer")
	}

	conf := config.LoadConfiguration("config/kafka.json")
	fmt.Println(conf)
	go pushData(producer, conf.TopicName)

	consumer, err := getKafkafkaConsumer()
	if err != nil {
		log.Printf("Error creating Consumer")
	}

	consumer.PullHandler(func(conMessage kafka.ConsumerMessage) {

		log.Println(conMessage.Message)
		log.Println(conMessage.Partition)
		log.Println(conMessage.Topic)
		log.Println(conMessage.Offset)
	})

}

func pushData(producer kafka.ProducerService, topic string) {
	i := 0
	for {
		time.Sleep(time.Second * 2)
		i = i + 1
		producer.Push(topic, getMessage(string(i)))
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type payload struct {
	Name string
	Age  int
}

func getMessage(name string) string {
	p := new(payload)
	p.Name = name
	p.Age = 33

	payloadJson, _ := json.Marshal(p)
	return string(payloadJson)
}
