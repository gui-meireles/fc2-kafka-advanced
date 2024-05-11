package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "fc2-kafka-advanced-kafka-1:9092",
		"client.id":         "goapp-consumer", // Id do consumer que esta lendo a mensagem do tópico / Caso troque o nome do consumer e mantenha o mesmo grupo, haverá um rebalance e os consumers do grupo irão ler partições diferentes
		"group.id":          "goapp-group",    // Inscreve o consumer em um grupo / Caso mude o nome do grupo, o consumer desse novo grupo, se for apenas 1, vai ler todas as partições
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
