package model

import "time"

// SongRepo -
type SongRepo struct {
	ID          int
	GroupName   *string `json:"group,omitempty"`
	SongName    *string `json:"song,omitempty"`
	ReleaseDate *string `json:"release_date"`
	Text        *string `json:"text"`
	Link        *string `json:"link"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
