package vo

type TopicVo struct {
	Id          int          `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	CoverImage  string       `json:"cover"`
	User        SimpleUserVo `json:"user"`
	TimeStamp   int64        `json:"timeStamp"`
	//CreateAt    int64        `json:"dateStr"`
}

type SimpleTopicVo struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	CoverImage  string `json:"cover,omitempty"`
}

//tTime.Format("2006年01月02日 03:04"))
