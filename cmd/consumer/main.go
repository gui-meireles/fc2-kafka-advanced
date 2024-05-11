package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "fc2-kafka-advanced-kafka-1:9092",
		"client.id":         "goapp-consumer", // Nome de quem está consumindo a mensagem do tópico
		"group.id":          "goapp-group",    // Inscreve o consumer em um grupo
		"auto.offset.reset": "earliest",       // Lê as mensagens desde o início
	}
	c, err := kafka.NewConsumer(configMap)
	if err != nil {
		fmt.Println("error consumer", err.Error())
	}
	topics := []string{"teste"}
	c.SubscribeTopics(topics, nil) // Inscreve o consumer no tópico escolhido = "teste"

	for {
		msg, err := c.ReadMessage(-1) // Timeout -1 = Ficará lendo infinitamente
		if err == nil {
			fmt.Println(string(msg.Value), msg.TopicPartition)
		}
	}
}
