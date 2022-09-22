package mq

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"go-micro.dev/v4/logger"
	"log"
	"vientiane/pub/adapter"
	pub "vientiane/pub/idl/grpc"
	"vientiane/server/models"
	"vientiane/server/service"
)

var (
	articleService = service.NewArticleService()
)

func ArticleConsumer() {
	// make a new reader that consumes from topic-A
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  "consumer-group-article",
		Topic:    adapter.TopicArticle,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	ctx := context.Background()
	for {
		m, err := r.FetchMessage(ctx)
		if err != nil {
			break
		}

		msg := &pub.ArticleMsg{}
		_ = json.Unmarshal(m.Value, msg)
		log.Println(msg)
		switch msg.WriteType {
		case adapter.ArticleTypeInsert:
			err = articleService.Add(ctx, &models.Article{
				Author:  msg.Article.Author,
				Title:   msg.Article.Title,
				Content: msg.Article.Content,
			})
		case adapter.ArticleTypeUpdate:
			err = articleService.Update(ctx, &models.Article{
				Id: msg.Article.Id,
				Author:  msg.Article.Author,
				Title:   msg.Article.Title,
				Content: msg.Article.Content,
			})
		case adapter.ArticleTypeDelete:
			err = articleService.Del(ctx, msg.Article.Id)
		}
		if err != nil {
			logger.Warnf("consumer article msg err: %+v",err)
			continue
		}

		if err = r.CommitMessages(ctx, m); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
