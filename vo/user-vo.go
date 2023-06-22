package vo

type UserVo struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Role     RoleVo `json:"role"`
}

type SimpleUserVo struct {
	Id       int    `gorm:"json:id"`
	Nickname string `gorm:"json:nickName"`
}
