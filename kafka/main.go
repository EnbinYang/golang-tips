package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	reader *kafka.Reader
	topic  = "user_click"
)

func writeKafka(ctx context.Context) {
	writer := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"), // broker addr
		Topic:                  topic,
		Balancer:               &kafka.Hash{},
		WriteTimeout:           1 * time.Second,
		RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: true,
	}
	defer writer.Close()

	// max retry three times
	for i := 0; i < 3; i++ {
		if err := writer.WriteMessages(
			ctx,
			kafka.Message{Key: []byte("1"), Value: []byte("enbin")},
			kafka.Message{Key: []byte("2"), Value: []byte("yang")},
			kafka.Message{Key: []byte("3"), Value: []byte("shaun")},
			kafka.Message{Key: []byte("1"), Value: []byte("hugo")},
			kafka.Message{Key: []byte("1"), Value: []byte("robot")},
		); err != nil {
			if err == kafka.LeaderNotAvailable {
				log.Println("Leader not available!")
				time.Sleep(500 * time.Millisecond)
				continue
			} else {
				log.Println("Error writing Kafka:", err)
			}
		} else {
			log.Println("Writing Kafka successful! Topic:", topic)
			break
		}
	}
}

func readKafka(ctx context.Context) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"localhost:9092"},
		Topic:          topic,
		CommitInterval: 1 * time.Second,
		GroupID:        "ai_team",
		StartOffset:    kafka.FirstOffset,
	})

	for {
		if msg, err := reader.ReadMessage(ctx); err != nil {
			log.Println("Error reading Kafka:", err)
			break
		} else {
			log.Printf("topic=%s, partition=%d, offset=%d, key=%s, value=%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
			time.Sleep(1 * time.Second)
		}
	}
}

func listenSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM) // ctrl+c && kill signal
	sig := <-c
	log.Printf("Catch signal %s", sig.String())
	if reader != nil {
		reader.Close()
	}
	os.Exit(0)
}

func main() {
	ctx := context.Background()
	writeKafka(ctx)
	readKafka(ctx)
}
