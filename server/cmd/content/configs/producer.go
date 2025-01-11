package configs

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/IBM/sarama"
	"github.com/google/uuid"
)

type KafkaProducer struct {
	topic string
			producer sarama.SyncProducer
}

func NewKafkaProducer(brokers string, topic string) (*KafkaProducer, error) {
	_brokers := strings.Split(brokers, ",")
	
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(_brokers, config)
	if err != nil {
		return nil, err
	}

	return &KafkaProducer{
		topic: topic,
		producer: producer,	
	},nil
}

func (p *KafkaProducer) SendMessage(value interface{}) error {
	json_data ,err := json.Marshal(value)
	if err != nil {
		return err
	}

	uuid := uuid.New()
	msg := &sarama.ProducerMessage{
		Topic: p.topic,
		Key: sarama.StringEncoder(uuid.String()),
		Value: sarama.StringEncoder(json_data),
	}

	partition, offset, err := p.producer.SendMessage(msg)
	
	if err != nil {
		log.Printf("Message send to partition %d at offset %d", partition, offset)
		return err
	}

	return nil
}
	

func (p *KafkaProducer) Close() error {
	return p.producer.Close()
}