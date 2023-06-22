package vo

type BlogVo struct {
	Id          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"desc"`
	CoverImage  string       `json:"coverImage"`
	DateStr     string       `json:"dateStr"`
	User        SimpleUserVo `json:"user"`
	Category    CategoryVo   `json:"category"`
}

type SimpleBlogVo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"desc,omitempty"`
	CoverImage  string `json:"coverImage,omitempty"`
	DateStr     string `json:"create,omitempty"`
}
