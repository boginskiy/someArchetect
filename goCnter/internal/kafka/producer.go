package kafka

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Mess struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	NewVal string `json:"new_val"`
}

func ProduceMess(message Mess) error {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		return err
	}
	defer p.Close()

	jsonMsg, _ := json.Marshal(message)
	topic := "counter-updates"
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          jsonMsg,
	}, nil)

	if err != nil {
		return err
	}

	return nil
}
