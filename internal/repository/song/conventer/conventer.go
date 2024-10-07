package conventer

import (
	"EffectiveMobile/internal/model"
	repoModel "EffectiveMobile/internal/repository/song/model"
)

// ToSongDetailFromSongRepo converting SongRepo object to SongDetailResponse object
func ToSongDetailFromSongRepo(info *repoModel.SongRepo) *model.SongDetailResponse {
	var song model.SongDetailResponse
	if info.GroupName != nil {
		song.GroupName = *info.GroupName
	}
	if info.SongName != nil {
		song.SongName = *info.SongName
	}
	if info.ReleaseDate != nil {
		song.ReleaseDate = *info.ReleaseDate
	}
	if info.Text != nil {
		song.Text = *info.Text
	}
	if info.Link != nil {
		song.Link = *info.Link
	}
	return &song
}
