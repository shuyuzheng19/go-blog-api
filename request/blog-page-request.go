package request

type BlogPageRequest struct {
	Page int  `json:"page"`
	Cid  int  `json:"cid"`
	Sort Sort `json:"sort"`
}

const (
	CREATE = "CREATE"
	UPDATE = "UPDATE"
	EYE    = "EYE"
	LIKE   = "LIKE"
	BACK   = "BACK"
)

type Sort string

func (s Sort) String() string {
	switch s {
	case CREATE:
		return "create_at desc"
	case UPDATE:
		return "update_at desc"
	case EYE:
		return "eye_count desc"
	case LIKE:
		return "like_count desc"
	case BACK:
		return "create_at asc"
	default:
		return "create_at desc"
	}
}
