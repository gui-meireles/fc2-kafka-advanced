package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func main() {
	deliveryChan := make(chan kafka.Event)
	producer := NewKafkaProducer()
	Publish("Mensagem", "teste", producer, nil, deliveryChan) // Publica a mensagem no tópico

	go DeliveryReport(deliveryChan) // Utilizando o 'go', jogamos essa função para outra thread -> Assíncrono

	fmt.Println("Guilherme Meireles")
	producer.Flush(1000)
}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "fc2-kafka-advanced-kafka-1:9092",
	}
	p, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}
	return p
}

func Publish(msg string, topic string, producer *kafka.Producer, key []byte, deliveryChannel chan kafka.Event) error {
	message := &kafka.Message{
		Value:          []byte(msg),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
	}
	err := producer.Produce(message, deliveryChannel)
	if err != nil {
		return err
	}
	return nil
}

func DeliveryReport(deliveryChan chan kafka.Event) { // Função para receber a resposta do evento enviado ao Kafka
	for e := range deliveryChan { // Para cada evento que ele receber, vai cair no loop e printar a mensagem
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Erro ao enviar")
			} else {
				fmt.Println("Mensagem enviada:", ev.TopicPartition)
				// Anotar no db que a mensagem foi processada
			}
		}
	}
}
