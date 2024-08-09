package entity

import "time"

type User struct {
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
}
