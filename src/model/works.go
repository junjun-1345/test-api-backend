package model

import (
	"time"
)

type Works struct {
	ID          uint      `gorm:"primarykey"`
	UserID      string    `gorm:"type:varchar(255)"`
	Date        int       `gorm:"type:int"`
	ClockIn     time.Time `gorm:"type:datetime"`
	ClockOut    time.Time `gorm:"type:datetime"`
	WorkingTime int64     `gorm:"type:int"`
}
