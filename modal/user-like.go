package modal

type UserLike struct {
	User   User   `json:"-"`
	UserId int    `gorm:"ForeignKey:usersId;unique"`
	Liked  string `gorm:"type:json;"`
}

type Like struct {
	Score  float64 `json:"Score"`
	Member string  `json:"Member"`
}
