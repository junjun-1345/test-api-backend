package model

import (
	"time"
)

type Shift struct {
	ID       uint      `gorm:"primarykey"`
	UserID   string    `gorm:"type:varchar(255)"`
	Date     int       `gorm:"type:int"`
	ClockIn  time.Time `gorm:"type:datetime"`
	ClockOut time.Time `gorm:"type:datetime"`
	// HACK
	Decide bool `gorm:"type:bool"`
}

type Shifts []Shift
