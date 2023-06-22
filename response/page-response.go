package response

type PageInfo struct {
	Page  int         `json:"page"`
	Size  int         `json:"size"`
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}
