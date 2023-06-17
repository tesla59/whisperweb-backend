package models

import "time"

type Confession struct {
	ID        string    `json:"id"`
	By        string    `json:"by"`
	To        string    `json:"to"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}