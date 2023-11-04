package request

import "time"

type ShiftRequest struct {
	UserID   string    `validate:"required" json:"userId"`
	Date     int       `validate:"required" json:"date"`
	ClockIn  time.Time `json:"clockIn"`
	ClockOut time.Time `json:"clockOut"`
	// HACK
	Decide bool `json:"decide"`
}

type CreateShiftsRequest struct {
	Shifts []ShiftRequest `json:"Shifts"`
}
