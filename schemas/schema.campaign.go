package schemas

import "time"

type Campaign struct {
	ID uint
	UserID uint
	Title string
	ShortDescription string
	Description string
	Slug string
	BackerCount int
	Perks string
	GoalAmount int
	CurrentAmount int
	CreatedAt time.Time
	UpdatedAt time.Time
}