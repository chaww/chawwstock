package model

import "time"

type Demo struct {
	ID        uint32    `json:"id"; gorm:"primary_key"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"time"`
}
