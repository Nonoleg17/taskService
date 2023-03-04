package entity

import "time"

type Session struct {
	Username string        `json:"username"`
	Value    string        `json:"value"`
	Expire   time.Duration `json:"expire"`
}
