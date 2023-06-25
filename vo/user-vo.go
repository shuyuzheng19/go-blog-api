package vo

type UserVo struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Icon     string `json:"icon"`
	Nickname string `json:"nickName"`
}

type SimpleUserVo struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickName"`
}

type CommentUserVo struct {
	Username string   `json:"username"`
	Avatar   string   `json:"avatar"`
	Level    int      `json:"level"`
	HomeLink string   `json:"homeLink"`
	LikeIds  []string `json:"likeIds"`
}
