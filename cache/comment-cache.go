package cache

type CommentCache interface {
	UserIsLikeComment(uid int, cid int64) bool
	AddUserLike(uid int, cid int64) int64
	CancelUserLike(uid int, cid int64) int64
	GetUserCommentLikes(uid int) []string
}
