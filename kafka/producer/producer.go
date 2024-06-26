package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"time"
)

func  main() {
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "topic_test", 0)
	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	conn.WriteMessages(kafka.Message{Value: []byte("Hello Kafka again")})

}
