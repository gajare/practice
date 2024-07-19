package main

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/lib/pq"
)

type Message struct {
	UserID  string `json:"user_id"`
	Content string `json:"content"`
}

func main() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "social_media_group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}
	defer consumer.Close()

	consumer.Subscribe("social_media", nil)

	db, err := sql.Open("postgres", "host=localhost port=5432 user=user password=password dbname=social_media sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}
	defer db.Close()

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			var message Message
			if err := json.Unmarshal(msg.Value, &message); err != nil {
				log.Printf("Failed to unmarshal message: %s", err)
				continue
			}
			_, err = db.Exec("INSERT INTO posts (user_id, content) VALUES ($1, $2)", message.UserID, message.Content)
			if err != nil {
				log.Printf("Failed to insert into database: %s", err)
			}
		} else {
			log.Printf("Consumer error: %s (%v)\n", err, msg)
		}
	}
}
