package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func main() {
	deliveryChan := make(chan kafka.Event)
	producer := NewKafkaProducer()
	Publish("Transferiu", "teste", producer, []byte("transferência"), deliveryChan) // Publica a mensagem no tópico / Ao colocar a key = []byte("transferência"), garantimos a ordem de recebimento das mensagens e que ela sempre irá apenas para uma partição

	go DeliveryReport(deliveryChan) // Utilizando o 'go', jogamos essa função para outra thread -> async

	producer.Flush(5000)
}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   "fc2-kafka-advanced-kafka-1:9092",
		"delivery.timeout.ms": "0",    // Tempo máx de entrega de uma mensagem ( 0 é infinito )
		"acks":                "all",  // 0 = Não tenho o retorno se a mensagem foi entregue (mais performatica) / 1 = Leader retorna que ele persistiu a mensagem / all = Retorna depois do Leader e os sync-brokers terem persistido a mensagem (menos performatica)
		"enable.idempotence":  "true", // Padrão é false = mensagem pode chegar repetida, pode perder alguma mensagem / true = Mensagem foi entregue na ordem e apenas 1 vez, caso utilize true, o "acks" deve ser all
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
