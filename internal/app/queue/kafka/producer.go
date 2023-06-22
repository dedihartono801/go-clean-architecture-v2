package kafka

import (
	"log"

	"github.com/Shopify/sarama"
)

type Producer interface {
	SendMessage(topic string, kafkaMessage string) error
}

type producer struct {
	producer sarama.SyncProducer
}

func NewKafkaProducer(broker string) (Producer, error) {
	// Set up configuration for the Kafka producer
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	// Create a new Kafka producer
	prd, err := sarama.NewSyncProducer([]string{broker}, config)
	if err != nil {
		return nil, err
	}

	return producer{prd}, nil

}

func (prd producer) SendMessage(topic string, kafkaMessage string) error {
	// Create a new Kafka message
	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(kafkaMessage),
	}

	// Send the message to Kafka
	partition, offset, err := prd.producer.SendMessage(message)
	if err != nil {
		return err
	}

	log.Printf("Message sent to partition %d at offset %d", partition, offset)
	return nil

}
