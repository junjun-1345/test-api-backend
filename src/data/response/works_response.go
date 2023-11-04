package response

import "time"

type WorksReponce struct {
	ID          uint      `json:"id"`
	UserID      string    `json:"userId"`
	Date        int       `json:"date"`
	ClockIn     time.Time `json:"clockIn"`
	ClockOut    time.Time `json:"clockOut"`
	WorkingTime int64     `json:"workingTime"`
}
