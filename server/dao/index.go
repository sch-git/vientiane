package dao

import (
	"context"
	"errors"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"strings"
)

type IndexDAO interface {
	CreateIndex(ctx context.Context, indexName, config string) (err error)
}

type indexDAO struct {
	esClient *elasticsearch.Client
}

func (i *indexDAO) CreateIndex(ctx context.Context, indexName, config string) (err error) {
	fun := "indexDAO.CreateIndex -->"
	if indexName == "" || config == "" {
		return errors.New(fun + "req is nil")
	}
	esResp, err := i.esClient.Indices.Create(
		"book_info",
		i.esClient.Indices.Create.WithPretty(),
		i.esClient.Indices.Create.WithBody(strings.NewReader(config)),
	)
	if err != nil {
		return
	}
	defer esResp.Body.Close()
	return
}

func NewIndexDAO() IndexDAO {
	cfg := elasticsearch.Config{
		Addresses: []string{"https://localhost:9200"},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	return &indexDAO{
		esClient: es,
	}
}
