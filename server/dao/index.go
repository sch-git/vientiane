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

	getResp, err := i.esClient.Indices.Exists([]string{indexName})
	if err != nil {
		return errors.New("index exists")
	}
	defer getResp.Body.Close()

	esResp, err := i.esClient.Indices.Create(
		indexName,
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
		Addresses: []string{"http://localhost:9200"},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("new es client -->")
	return &indexDAO{
		esClient: es,
	}
}
