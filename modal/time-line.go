package modal

type TimeLine struct {
	Id      int       `json:"-";gorm:"primary_key"`
	Title   string    `json:"title";gorm:'size:100;default:""'`
	Content string    `json:"content"`
	Time    LocalTime `json:"time"`
	UserId  int       `json:"-"`
	User    User      `json:"-"`
}
