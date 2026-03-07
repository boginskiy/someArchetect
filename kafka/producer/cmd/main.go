package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal // Ждать подтверждения хотя бы от лидера партиции
	config.Producer.Retry.Max = 5                      // Максимум попыток отправить сообщение
	config.Producer.Return.Successes = true            // Возвращать успешные отправки

	brokers := []string{"localhost:9092"} // Список брокеров Kafka

	// Создаем синхронного продюсера
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		panic(err)
	}

	defer producer.Close()

	// Создаем сообщение
	msg := &sarama.ProducerMessage{
		Topic: "my-topic",                                  // Название темы
		Value: sarama.StringEncoder("Hello Kafka from Go"), // Данные сообщения
	}

	// Отправляем сообщение
	partition, offset, err := producer.SendMessage(msg) // Отправляем сообщение
	if err != nil {
		fmt.Println("Error sending message:", err)
		os.Exit(1)
	}

	fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)

	// Graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	sig := <-signals
	fmt.Printf("Caught signal %v: terminating\n", sig)
}
