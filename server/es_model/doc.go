package es_model

type ESSearchResp struct {
	Took    int64     `json:"took"`
	TimeOut bool      `json:"time_out"`
	Hits    *HitsResp `json:"hits"`
}

type HitsResp struct {
	Total *TotalInfo  `json:"total"`
	Hits  []*HitsInfo `json:"hits"`
}

type TotalInfo struct {
	Value    int64  `json:"value"`
	Relation string `json:"relation"`
}

type HitsInfo struct {
	Index  string                 `json:"_index"`
	Id     string                 `json:"_id"`
	Source map[string]interface{} `json:"_source"`
	Sort   []interface{}          `json:"sort"`
}
