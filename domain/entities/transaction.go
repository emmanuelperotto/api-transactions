package entities

import "time"

type Transaction struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	AmountCents int64     `json:"amountCents"`
	EventDate   time.Time `json:"eventDate"`
	AccountID   int
}
