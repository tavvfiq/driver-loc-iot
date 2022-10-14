package mykafka

import (
	"context"

	kfk "github.com/segmentio/kafka-go"
)

type KafkaWriter interface {
	Push(ctx context.Context, message kfk.Message) error
}

type KafkaReader interface {
	Read(ctx context.Context) (kfk.Message, error)
	Close() error
}

type service struct {
	writer *kfk.Writer
	reader *kfk.Reader
}
