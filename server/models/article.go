package models

import (
	"context"
	pub "vientiane/pub/idl/grpc"
)

const articleTableName = "vientiane_article"

type ArticleService interface {
	Get(ctx context.Context, id int64) (*Article, error)
	List(ctx context.Context, article *Article) ([]*Article, error)
	Add(ctx context.Context, article *Article) (err error)
	Del(ctx context.Context, article *Article) (err error)
}

type Article struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
	Limit   int64  `json:"limit" gorm:"-"`
	Offset  int64  `json:"offset" gorm:"-"`
}

func (m *Article) ToGrpc() *pub.Article {
	article := &pub.Article{}
	if nil == m || m.IsEmpty() {
		return article
	}

	return &pub.Article{
		Id:      m.Id,
		Title:   m.Title,
		Content: m.Content,
		Author:  m.Author,
	}
}

func (m *Article) TableName() string {
	return articleTableName
}

func (m Article) IsEmpty() bool {
	return m == Article{}
}
