package models

import "time"

type Meeting struct {
	ID              int    `gorm:"not null"`
	BrigadeID       int `gorm:"not null"`
	Brigade         Brigade
	Date            time.Time `gorm:"not null"`
	AttendanceCount int       `gorm:"not null"`
}
