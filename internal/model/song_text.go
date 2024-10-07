package model

// SongText -
type SongText struct {
	Text []Couplet `json:"text"`
}

// Couplet -
type Couplet struct {
	Couplet int      `json:"couplet"`
	Text    []string `json:"text"`
}
