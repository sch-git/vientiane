package test

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"testing"
)

var (
	esCli *elasticsearch.Client
	ctx   = context.Background()
)

func init() {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	esCli, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
}

func TestIndex(t *testing.T) {
	esResp, err := esCli.Info()
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
	defer esResp.Body.Close()

	t.Log(esResp)
}
