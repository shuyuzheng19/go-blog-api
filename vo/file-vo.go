package vo

type FileVo struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	DateStr string `json:"dateStr"`
	Suffix  string `json:"suffix"`
	SizeStr string `json:"sizeStr"`
	MD5     string `json:"md5"`
	URL     string `json:"url"`
}
