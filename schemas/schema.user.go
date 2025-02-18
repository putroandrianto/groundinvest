package schemas

import "time"

type User struct {
	ID uint
	Name string
	Occupation string
	Email string
	PasswordHash string
	AvatarFileName 	string
	Role string
	Token string
	CreatedAt time.Time
	UpdatedAt time.Time
}