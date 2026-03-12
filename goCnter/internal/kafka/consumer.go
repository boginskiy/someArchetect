package kafka

import (
	"database/sql"
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ConsumeMessagesFromKafka(c *kafka.Consumer, topic string, dbConn *sql.DB) {
	c.SubscribeTopics([]string{topic}, nil)

	for {
		msg, err := c.ReadMessage(3 * time.Second)
		if err == nil {
			var m Message
			err := json.Unmarshal(msg.Value, &m)
			if err != nil {
				log.Printf("Ошибка разбора JSON: %s", err)
				continue
			}

			switch m.Type {

			case "increment":
				count, err := strconv.Atoi(m.NewVal)
				if err != nil {
					log.Printf("Ошибка преобразования строки в число: %s", err)
					break
				}
				err = db.UpdateCounter(dbConn, m.ID, count)
				if err != nil {
					log.Printf("Ошибка обновления базы данных: %s", err)
				}

			default:
				log.Printf("Неизвестный тип сообщения: %s", m.Type)
			}
		} else {
			log.Printf("Ошибка чтения сообщения: %s", err)
		}
	}
}
