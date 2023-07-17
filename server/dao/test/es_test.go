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
	"vientiane/server/es_model"
	"vientiane/server/service/entity"
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
            "sort.field": "book_created_at",
            "sort.order": "desc"
        },
        "index.store.preload": [
            "nvd",
            "dvd",
            "tim",
            "tip",
            "doc"
        ]
    },
    "mappings": {
        "properties": {
            "book_id": {
                "type": "long"
            },
            "book_title": {
                "type": "keyword"
            },
            "client_id": {
                "type": "integer"
            },
            "book_category_list": {
                "type": "keyword"
            },
            "book_category_name_list": {
                "type": "keyword"
            },
            "book_status": {
                "type": "integer"
            },
            "author": {
                "type": "keyword"
            },
            "author_id": {
                "type": "long"
            },
            "account_id": {
                "type": "long"
            },
            "user_id": {
                "type": "long"
            },
            "editor": {
                "type": "keyword"
            },
            "editor_id": {
                "type": "long"
            },
            "total_chapter_num": {
                "type": "integer"
            },
            "book_created_at": {
                "format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis",
                "type": "date"
            },
            "book_updated_at": {
                "format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis",
                "type": "date"
            }
        }
    }
}
    `
	esResp, err := esCli.Indices.Create(
		"book_info",
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

// bulk upsert
func TestBulkInsert(t *testing.T) {
	infos := []*es_model.ESBookInfo{
		{BookId: 100001, BookTitle: "大爱如烟", ClientId: 0, AuthorId: 10011, AccountId: 31, BookCategoryList: []int64{1, 17}, BookCategoryNameList: []string{"现代言情", "总裁豪门"}, BookCreatedAt: "2017-04-14 11:23:59", BookUpdatedAt: "2023-06-01 02:30:00"},
		{BookId: 100002, BookTitle: "上海上海", ClientId: 1, AuthorId: 10016, AccountId: 36, BookCategoryList: []int64{8, 42}, BookCategoryNameList: []string{"都市人生", "都市高手"}, BookCreatedAt: "2017-04-14 11:39:29", BookUpdatedAt: "2023-06-01 02:30:00"},
		{BookId: 100003, BookTitle: "北京", ClientId: 1, AuthorId: 10017, AccountId: 37, BookCategoryList: []int64{8, 42}, BookCategoryNameList: []string{"都市人生", "都市高手"}, BookCreatedAt: "2017-04-14 11:39:30", BookUpdatedAt: "2023-06-01 02:30:00"},
		{BookId: 100004, BookTitle: "zHE", ClientId: 1, AuthorId: 10016, AccountId: 36, BookCategoryList: []int64{8, 42}, BookCategoryNameList: []string{"都市人生", "都市高手"}, BookCreatedAt: "2017-04-14 11:39:31", BookUpdatedAt: "2023-06-01 02:30:00"},
	}

	items := make([]*es_model.ESBulkItem, 0)
	for _, info := range infos {
		bs, _ := json.Marshal(info)
		items = append(items, &es_model.ESBulkItem{Index: "book_info", DocId: fmt.Sprintf("%d", info.BookId), DocData: []byte(fmt.Sprintf(`{"doc":%s, "doc_as_upsert": true}`, string(bs)))})
	}

	buf := buffer.Buffer{}
	for _, item := range items {
		buf.WriteString(fmt.Sprintf(`{"update":{"_index":"%s","_id":"%s"}}`+"\n", item.Index, item.DocId))
		buf.WriteString(string(item.DocData) + "\n")
	}

	log.Println(buf.String())
	esResp, err := esCli.Bulk(
		strings.NewReader(buf.String()),
		esCli.Bulk.WithIndex("book_info"),
		esCli.Bulk.WithPretty(),
	)
	if err != nil {
		t.Logf("err: %#v", err)
		return
	}
	defer esResp.Body.Close()
	log.Println(esResp)
}

func TestBulkDel(t *testing.T) {
	buf := buffer.Buffer{}
	buf.WriteString(fmt.Sprintf(`{"delete":{"_index":"%s", "_id":"%s"}}`+"\n", "book_info", "100001"))
	buf.WriteString(fmt.Sprintf(`{"delete":{"_index":"%s", "_id":"%s"}}`+"\n", "book_info", "100002"))
	log.Println(buf.String())

	esResp, err := esCli.Bulk(
		strings.NewReader(buf.String()),
		esCli.Bulk.WithIndex("book_info"),
		esCli.Bulk.WithPretty(),
	)
	if err != nil {
		t.Logf("err: %#v", err)
		return
	}
	defer esResp.Body.Close()
	log.Println(esResp)
}

// es util bulk
func TestESUtilBulkInsert(t *testing.T) {

}

// 检索
func TestSearch(t *testing.T) {
	//	query := `
	//{
	//	"_source": ["_id"],
	//    "query": {
	//        "bool": {
	//            "should": [
	//                {
	//                    "match": {
	//                        "client_id": 1
	//                    }
	//                },
	//                {
	//                    "match": {
	//                        "client_id": 4
	//                    }
	//                }
	//            ],
	//            "must": [
	//                {
	//                    "match": {
	//                        "client_id": 4
	//                    }
	//                }
	//            ]
	//        }
	//    }
	//}
	//`

	conditions := &entity.Conditions{
		//Musts: []entity.Conditions{
		//{Cond: &entity.Condition{Field: "field", OpType: "exists", Value: "book_category_list"}},
		//{Cond: &entity.Condition{Field: "book_category_list", OpType: consts.ESOpTypeIn, Value: []int64{8, 1}}},
		//{Cond: &entity.Condition{Field: "book_id", OpType: consts.ESOpTypeEq, Value: 100001}},
		//{Cond: &entity.Condition{Field: "book_title", OpType: consts.ESOpTypeEq, Value: "大爱如烟"}},
		//},
		//Should: []entity.Conditions{
		//	{
		//		Musts: []entity.Conditions{
		//			{Cond: &entity.Condition{Field: "book_id", OpType: consts.ESOpTypeEq, Value: "100001"}},
		//			{Cond: &entity.Condition{Field: "book_title", OpType: consts.ESOpTypeEq, Value: "大爱如烟"}},
		//		},
		//	},
		//	{
		//		Cond: &entity.Condition{Field: "book_id", OpType: consts.ESOpTypeEq, Value: "100002"},
		//	},
		//},
		//MustNot: []entity.Conditions{
		//	{Cond: &entity.Condition{Field: "book_id", OpType: consts.ESOpTypeIn, Value: []string{"100001"}}},
		//},
		//MustNot: []entity.Conditions{{Cond: &entity.Condition{Filed: "field", OpType: "exists", Value: "book_category_list"}}},
		//Cond: &entity.Condition{Field: "client_id", OpType: "=", Value: 1},
	}

	query := entity.ParseToES(conditions)
	queryParam := &entity.ESQueryParam{
		Source: []string{"book_id", "book_title"},
		Sort:   []string{"author_id:asc", "book_title:desc"},
		Query:  query,
	}
	queryBytes, _ := json.Marshal(map[string]interface{}{"query": queryParam.Query})
	//log.Println(string(queryBytes))
	esResp, err := esCli.Search(
		esCli.Search.WithIndex("book_info"),
		esCli.Search.WithSize(10),
		esCli.Search.WithFrom(0),
		esCli.Search.WithBody(strings.NewReader(string(queryBytes))),
		esCli.Search.WithSort(queryParam.Sort...),
		esCli.Search.WithSource(queryParam.Source...),
	)
	if err != nil {
		t.Errorf("err:%#v", err)
		return
	}
	defer esResp.Body.Close()

	queryResp := &es_model.ESSearchResp{}
	err = json.NewDecoder(esResp.Body).Decode(queryResp)
	if err != nil {
		t.Errorf("err:%+v", err)
		return
	}

	t.Logf("%+v", queryResp)
}

func TestExists(t *testing.T) {
	esResp, err := esCli.Indices.Exists([]string{"book_info", "t_info", "xx"})
	if err != nil {
		t.Errorf("err: %+v", err)
		return
	}

	var m = make(map[string]interface{})
	json.NewDecoder(esResp.Body).Decode(&m)

	t.Logf("resp: %+v", esResp)
	t.Logf("body: %+v", m)
	t.Logf("code-200: %+v", esResp.StatusCode == 200)
	t.Logf("code-404: %+v", esResp.StatusCode == 404)
}
