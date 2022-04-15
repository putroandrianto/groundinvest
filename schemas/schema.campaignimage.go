package schemas

import "time"

type CampaignImage struct {
	ID uint
	CampaignID uint
	Filename string
	is_primary bool
	CreatedAt time.Time
	UpdatedAt time.Time
}