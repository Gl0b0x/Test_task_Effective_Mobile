package model

// SongDetailResponse -
type SongDetailResponse struct {
	GroupName   string `json:"group" example:"Muse"`
	SongName    string `json:"song" example:"Supermassive Black Hole"`
	ReleaseDate string `json:"release_date" example:"14.03.2024"`
	Text        string `json:"text" example:"Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight"`
	Link        string `json:"link" example:"https://www.youtube.com/watch?v=Xsp3_a-PMTw"`
}

// SongDetailRequest -
type SongDetailRequest struct {
	GroupName   *string `json:"group" example:"Muse"`
	SongName    *string `json:"song" example:"Supermassive Black Hole"`
	ReleaseDate *string `json:"release_date" example:"14.03.2024" omitempty:"true"`
	Text        *string `json:"text" omitempty:"true" example:"Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight"`
	Link        *string `json:"link" omitempty:"true" example:"https://www.youtube.com/watch?v=Xsp3_a-PMTw"`
}
