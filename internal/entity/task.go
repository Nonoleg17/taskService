package entity

import "time"

type Task struct {
	ID          int64     `json:"id"`
	Header      string    `json:"header"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
	Creation    time.Time `json:"creation"`
}
