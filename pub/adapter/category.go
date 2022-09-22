package adapter

type Category struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

const (
	ArticleTypeInsert      = "insert"
	ArticleTypeUpdate      = "update"
	ArticleTypeDelete      = "delete"
	TopicArticle           = "article"
	TopicArticlePartition0 = 0
)
