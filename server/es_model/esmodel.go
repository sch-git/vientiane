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

type ESBookInfo struct {
	BookId               int64    `json:"book_id"`
	BookTitle            string   `json:"book_title"`
	ClientId             int64    `json:"client_id"`
	BookCategoryList     []int64  `json:"book_category_list"`
	BookCategoryNameList []string `json:"book_category_name_list"`
	BookStatus           int64    `json:"book_status"`
	Author               string   `json:"author"`
	AuthorId             int64    `json:"author_id"`
	AccountId            int64    `json:"account_id"`
	UserId               int64    `json:"user_id"`
	Editor               string   `json:"editor"`
	EditorId             int64    `json:"editor_id"`
	TotalChapterNum      int64    `json:"total_chapter_num"`
	BookCreatedAt        string   `json:"book_created_at"`
	BookUpdatedAt        string   `json:"book_updated_at"`
}
