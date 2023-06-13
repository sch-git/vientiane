package es_model

type ESBulkItem struct {
	Index   string `json:"index"`
	DocId   string `json:"doc_id"`
	DocData []byte `json:"doc_data"`
}

type ESAccountInfo struct {
	AuthorIds []int64 `json:"author_ids"`
	AccountId int64   `json:"account_id"`
	UserId    int64   `json:"user_id"`
	ClientId  []int64 `json:"client_id"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
