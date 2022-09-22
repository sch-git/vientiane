package utils

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	"vientiane/pub/adapter"
	pub "vientiane/pub/idl/grpc"
)

var (
	conn *kafka.Conn
	err error
)

func KafkaClose()  {
	conn.Close()
}

func init()  {
	// to produce messages
	conn, err = kafka.DialLeader(context.Background(), "tcp", "localhost:9092", adapter.TopicArticle, adapter.TopicArticlePartition0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	//_ = conn.SetWriteDeadline(time.Now().Add(10*time.Second))
}

func WriteMsg(msg *pub.ArticleMsg)  {
	msgBytes,err := json.Marshal(msg)
	_, err = conn.WriteMessages(
		kafka.Message{Value: msgBytes},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
}
