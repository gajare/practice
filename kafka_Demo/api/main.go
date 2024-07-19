package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gorilla/mux"
)

type Message struct {
	UserID  string `json:"user_id"`
	Content string `json:"content"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/post", handlePost).Methods("POST")

	http.ListenAndServe(":8080", r)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	var msg Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		http.Error(w, "Failed to create producer", http.StatusInternalServerError)
		return
	}
	defer p.Close()

	topic := "social_media"
	value, _ := json.Marshal(msg)
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          value,
	}, nil)
	if err != nil {
		http.Error(w, "Failed to produce message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "Message accepted")
}
