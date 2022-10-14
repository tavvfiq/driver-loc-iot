package mykafka

import (
	"context"
	"time"

	kfk "github.com/segmentio/kafka-go"
)

func NewReader(urls []string, clientId, topic string, timeout time.Duration) KafkaReader {
	config := kfk.ReaderConfig{
		Brokers:         urls,
		GroupID:         clientId,
		Topic:           topic,
		MinBytes:        10e3,
		MaxBytes:        10e6,
		MaxWait:         timeout,
		ReadLagInterval: -1,
	}
	return &service{
		reader: kfk.NewReader(config),
	}
}

func (s *service) Read(ctx context.Context) (kfk.Message, error) {
	return s.reader.ReadMessage(ctx)
}

func (s *service) Close() error {
	return s.reader.Close()
}
