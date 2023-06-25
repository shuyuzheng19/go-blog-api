package search

type SearchResponse struct {
	Hits               interface{} `json:"hits"`
	Offset             int         `json:"offset"`
	Limit              int         `json:"limit"`
	EstimatedTotalHits int64       `json:"estimatedTotalHits"`
	TotalHits          int64       `json:"totalHits"`
	TotalPages         int         `json:"totalPages"`
	HitsPerPage        int         `json:"hitsPerPage"`
	Page               int         `json:"page"`
	QrocessingTimeMs   int         `json:"qrocessingTimeMs"`
	Query              string      `json:"query"`
}
