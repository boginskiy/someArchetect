package main

import (
	"context"
	"fmt"
	"kafka/consumer/handlers"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/IBM/sarama"
)

func main() {
	brokers := []string{"localhost:9092"}
	groupID := "my-consumer-group"
	topics := []string{"my-topic"}

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true // Возврат ошибок потребителей
	config.Version = sarama.V2_6_0_0     // Версия протокола Kafka

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		panic(err)
	}
	defer consumerGroup.Close()

	// Context
	ctx, cancel := context.WithCancel(context.Background())

	// Graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// WaitGroup
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			handler := &handlers.ConsumerGroupHandler{}
			err := consumerGroup.Consume(ctx, topics, handler)
			if err != nil {
				log.Printf("Error consuming messages: %v", err)
			}

			if ctx.Err() != nil {
				return
			}
		}
	}()

	sig := <-signals
	fmt.Printf("Caught signal %v: terminating\n", sig)

	cancel()
	wg.Wait()
}
