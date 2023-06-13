package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"go.uber.org/zap/buffer"
	"log"
	"strings"
	"testing"
	"time"
	"vientiane/server/es_model"
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

func TestInfo(t *testing.T) {
	esResp, err := esCli.Info()
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
	defer esResp.Body.Close()

	t.Log(esResp)
}

// 索引信息
func TestCatIndices(t *testing.T) {
	esResp, err := esCli.Cat.Indices(
		esCli.Cat.Indices.WithPretty(),
		esCli.Cat.Indices.WithV(true),
		esCli.Cat.Indices.WithS("index"),
		esCli.Cat.Indices.WithFormat("json"),
	)
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
	defer esResp.Body.Close()

	t.Log(esResp)
}

// 创建索引
func TestAddIndices(t *testing.T) {
	config := `
{
    "settings": {
        "number_of_shards": 2,
        "number_of_replicas": 0,
        "refresh_interval": "1s",
        "index": {
          "sort.field": "created_at",
          "sort.order": "desc"
        },
        "index.store.preload": ["nvd", "dvd", "tim", "tip", "doc"]
    },
    "mappings": {
        "properties": {
            "author_ids": {
                "type": "keyword"
            },
            "account_id": {
                "type": "keyword"
            },
            "user_id": {
                "type": "keyword"
            },
            "client_id": {
                "type": "integer"
            },
            "created_at": {
                "format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis",
                "type": "date"
            },
            "updated_at": {
                "format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis",
                "type": "date"
            }
        }
    }
}
    `
	esResp, err := esCli.Indices.Create(
		"api-index",
		esCli.Indices.Create.WithPretty(),
		esCli.Indices.Create.WithBody(strings.NewReader(config)),
	)
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
	defer esResp.Body.Close()

	t.Log(esResp)
}

// bulk 写入
func TestBulkInsert(t *testing.T) {
	infos := []*es_model.ESAccountInfo{
		{AuthorIds: []int64{1, 2, 3}, AccountId: 10, UserId: 10, ClientId: []int64{1, 2, 3, 4}, CreatedAt: time.Now().Format("2006-01-02 15:04:05"), UpdatedAt: time.Now().Format("2006-01-02 15:04:05")},
		{AuthorIds: []int64{2, 3, 4}, AccountId: 11, UserId: 11, ClientId: []int64{1, 2, 3}, CreatedAt: time.Now().Format("2006-01-02 15:04:05"), UpdatedAt: time.Now().Format("2006-01-02 15:04:05")},
	}

	items := make([]*es_model.ESBulkItem, 0)
	for _, info := range infos {
		bs, _ := json.Marshal(info)
		items = append(items, &es_model.ESBulkItem{Index: "api-index", DocId: fmt.Sprintf("%d", info.AccountId), DocData: []byte(fmt.Sprintf(`{"doc":%s, "doc_as_upsert": true}`, string(bs)))})
	}

	buf := buffer.Buffer{}
	for _, item := range items {
		buf.WriteString(fmt.Sprintf(`{"update":{"_index":"%s","_id":"%s"}}`+"\n", item.Index, item.DocId))
		buf.WriteString(string(item.DocData) + "\n")
	}

	log.Println(buf.String())
	esResp, err := esCli.Bulk(
		strings.NewReader(buf.String()),
		esCli.Bulk.WithIndex("api-index"),
		esCli.Bulk.WithPretty(),
	)
	if err != nil {
		t.Logf("err: %#v", err)
		return
	}
	defer esResp.Body.Close()
	log.Println(esResp)
}

// 检索
func TestSearch(t *testing.T) {
	query := `
{
	"_source": ["_id"],
    "query": {
        "bool": {
            "should": [
                {
                    "match": {
                        "client_id": 1
                    }
                },
                {
                    "match": {
                        "client_id": 4
                    }
                }
            ],
            "must": [
                {
                    "match": {
                        "client_id": 4
                    }
                }
            ]
        }
    }
}
`
	esResp, err := esCli.Search(
		esCli.Search.WithIndex("api-index"),
		esCli.Search.WithSize(10),
		esCli.Search.WithFrom(0),
		esCli.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		t.Errorf("err:%#v", err)
		return
	}
	defer esResp.Body.Close()

	t.Log(esResp)
}
