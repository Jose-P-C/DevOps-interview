package models

import "time"

type Course struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
