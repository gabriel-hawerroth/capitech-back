package entity

import "time"

type SearchLog struct {
	Id         int       `json:"id"`
	UserId     int       `json:"user_id"`
	FieldKey   string    `json:"field_key"`
	FieldValue string    `json:"field_value"`
	SearchDate time.Time `json:"search_date"`
}
