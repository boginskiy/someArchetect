package handlers

import (
	"fmt"

	"github.com/IBM/sarama"
)

type ConsumerGroupHandler struct{}

// Implement the setup method (optional)
func (h *ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

// Implement the cleanup method (optional)
func (h *ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

// The core logic goes here
func (h *ConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {

		fmt.Printf("Received message: topic:%s, partition:%d, offset:%d, key:%s, value:%s\n",
			msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))

		sess.MarkMessage(msg, "") // Mark the message as processed
	}

	return nil
}
