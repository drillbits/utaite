package utaite

import "time"

// Member represents a member of utaite.
type Member struct {
	Name      string
	Roles     []string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Role represents a role of member.
type Role struct {
	Name string
}
