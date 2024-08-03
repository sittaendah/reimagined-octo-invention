package mb

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

var Producer *kafka.Producer

type Payload struct {
	Type    string `json:"type"`
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func SetupKafkaProducer() {
	config := &kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BOOTSTRAP_SERVER"),
		"security.protocol": os.Getenv("KAFKA_SECURITY_PROTOCOL"),
		"sasl.mechanisms":   os.Getenv("KAFKA_SASL_MECHANISM"),
		"sasl.username":     os.Getenv("KAFKA_SASL_USERNAME"),
		"sasl.password":     os.Getenv("KAFKA_SASL_PASSWORD"),
	}

	var err error
	Producer, err = kafka.NewProducer(config)
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
	}
}

func SendMessage(username string) {
	topic := "aegis.hiring"
	payload := Payload{
		Type:    "login",
		Status:  true,
		Message: username + " logged in",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to serialize payload: %s", err)
		return
	}

	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          payloadBytes,
	}

	err = Producer.Produce(msg, nil)
	if err != nil {
		log.Printf("Failed to produce message: %s", err)
	}
	Producer.Flush(15 * 1000)
}
