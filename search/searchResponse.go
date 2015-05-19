package search

type SearchResponse struct {
	Took  int64      `json:"took"`
	Hits  SearchHits `json:"hits"`
	Error string     `json:"error"`
}

func (self *SearchResponse) GetHits() []SearchHit {
	return self.Hits.Hits
}

func (self *SearchResponse) TotalHits() int64 {
	return self.Hits.TotalHits
}

type SearchHits struct {
	TotalHits int64       `json:"total"`
	Hits      []SearchHit `json:"hits"`
}

type SearchHit struct {
	Score  float64                `json:"_score"`
	Index  string                 `json:"_index"`
	Id     string                 `json:"_id"`
	Type   string                 `json:"_type"`
	Fields map[string]interface{} `json:"fields"`
}
