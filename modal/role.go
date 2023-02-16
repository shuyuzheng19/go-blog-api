package modal

type Role struct {
	Id          int    `gorm:"primary_key"`
	Name        string `gorm:"not null;unique"`
	Description string `gorm:"not null"`
}
