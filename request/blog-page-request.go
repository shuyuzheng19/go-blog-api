package request

type BlogPageRequest struct {
	Page int  `json:"page" form:"page"`
	Cid  int  `json:"cid" form:"cid"`
	Sort Sort `json:"sort" form:"sort"`
}

func (request *BlogPageRequest) SetPageDefaultValue() {
	if request.Page <= 0 {
		request.Page = 1
	}
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
