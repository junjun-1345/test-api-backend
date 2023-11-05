package model

type User struct {
	ID             uint   `gorm:"primarykey"`
	UserID         string `gorm:"type:varchar(255)"`
	Name           string `gorm:"type:varchar(255)"`
	WorkInWeekDay  int    `gorm:"type:int"`
	WorkInWeekTime int    `gorm:"type:int"`
	Rank           int    `gorm:"type:int"`
	Vacation       int    `gorm:"type:int"`
	Admin          bool   `gorm:"type:bool"`
}
