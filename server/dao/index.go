package dao

import (
	"github.com/elastic/go-elasticsearch/v8"
	"log"
)

type IndexDAO interface {
}

type indexDAO struct {
	esClient *elasticsearch.Client
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
