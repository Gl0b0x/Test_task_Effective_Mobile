package repository

import (
	"EffectiveMobile/internal/model"
	"EffectiveMobile/internal/repository/song"
	"context"
	"database/sql"
)

// SongRepository interface repository
type SongRepository interface {
	CreateSong(ctx context.Context, info *model.SongDetailRequest) error
	DeleteSongByName(ctx context.Context, info *model.SongRequest) error
	GetAllSongs(ctx context.Context, limit int, offset int) ([]*model.SongDetailResponse, error)
	GetSongByName(ctx context.Context, info *model.SongRequest) (*model.SongDetailResponse, error)
	GetSongsByFilter(ctx context.Context, info *model.SongDetailRequest, limit int, offset int) ([]*model.SongDetailResponse, error)
	UpdateSongByName(ctx context.Context, req *model.SongRequest, info *model.SongDetailRequest) (*model.SongDetailResponse, error)
}

// Repository -
type Repository struct {
	SongRepository
}

// NewRepository -
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		SongRepository: song.NewSongPostgres(db),
	}
}
