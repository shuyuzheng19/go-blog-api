package vo

type CommentVo struct {
	ID         int64         `json:"id"`
	ParentId   *int64        `json:"parentId"`
	User       CommentUserVo `json:"user"`
	Address    string        `json:"address"`
	Content    string        `json:"content"`
	UID        int           `json:"uid"`
	Likes      int64         `json:"likes"`
	CreateTime string        `json:"createTime"`
	ContentImg string        `json:"contentImg"`
	Reply      ReplyVo       `json:"reply"`
}

type ReplyVo struct {
	Total int         `json:"total"`
	List  []CommentVo `json:"list,"`
}
