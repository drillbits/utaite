package utaite

import "time"

// Member represents a member of utaite.
type Member struct {
	Name      string    `json:"name"`
	Roles     []string  `json:"roles"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Twitter represents a member's twitter account
type Twitter struct {
	ScreenName      string    `json:"screenName"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	URL             string    `json:"url"`
	ProfileImageURL string    `json:"profileImageURL"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

// MemberTwitter represents an association between a member and a twitter account.
type MemberTwitter struct {
	MemberID  int64     `json:"memberID"`
	TwitterID int64     `json:"twitterID"`
	CreatedAt time.Time `json:"createdAt"`
}

// Role represents a role of member.
type Role struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Mylist represents a nicovideo's mylist of member.
type Mylist struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
