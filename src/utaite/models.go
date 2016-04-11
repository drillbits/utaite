package utaite

import "time"

// Member represents a member of utaite.
type Member struct {
	Name      string    `json:"name"`
	Roles     []string  `json:"roles"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Role represents a role of member.
type Role struct {
	Name string `json:"name"`
}
