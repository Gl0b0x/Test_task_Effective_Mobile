package service

import (
	"EffectiveMobile/internal/model"
	"EffectiveMobile/internal/repository"
	"EffectiveMobile/internal/service/song"
	"context"
)

// SongService -
type SongService interface {
	CreateSong(ctx context.Context, info *model.SongDetailRequest) (*model.SongDetailResponse, error)
	DeleteSongByName(ctx context.Context, info *model.SongRequest) (bool, error)
	GetAllSongs(ctx context.Context, limit int, offset int) ([]*model.SongDetailResponse, error)
	UpdateSongByName(ctx context.Context, req *model.SongRequest, info *model.SongDetailRequest) (*model.SongDetailResponse, error)
	GetSongByName(ctx context.Context, info *model.SongRequest) (*model.SongDetailResponse, error)
	GetSongText(ctx context.Context, info *model.SongRequest) (*model.SongText, error)
	GetSongsByFilter(ctx context.Context, info *model.SongDetailRequest, limit int, offset int) ([]*model.SongDetailResponse, error)
}

// Service -
type Service struct {
	SongService
}

// NewService -
func NewService(repo *repository.Repository) *Service {
	return &Service{
		SongService: song.NewService(repo),
	}
}
