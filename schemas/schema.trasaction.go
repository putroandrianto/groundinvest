package schemas

import "time"

type Transaction struct {
	ID uint
	UserID uint
	CampaignID uint
	Amount int
	Status string
	Code string
	CreatedAt time.Time
	UpdatedAt time.Time
}