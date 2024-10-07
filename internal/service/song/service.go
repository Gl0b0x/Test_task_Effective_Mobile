package song

import (
	"EffectiveMobile/internal/model"
	"EffectiveMobile/internal/repository"
	"context"
	"database/sql"
	"errors"
	"log"
	"strings"
)

type service struct {
	repo repository.SongRepository
}

// NewService create service
func NewService(repo *repository.Repository) *service {
	return &service{repo: repo}
}

func (s *service) UpdateSongByName(ctx context.Context, req *model.SongRequest, info *model.SongDetailRequest) (*model.SongDetailResponse, error) {
	song, err := s.repo.UpdateSongByName(ctx, req, info)
	if err != nil {
		return nil, err
	}
	return song, nil
}

func (s *service) CreateSong(ctx context.Context, info *model.SongDetailRequest) (*model.SongDetailResponse, error) {
	err := s.repo.CreateSong(ctx, info)
	if err != nil {
		log.Printf("ошибка создания песни: %v\n", err)
		return nil, err
	}
	songRequest := model.SongRequest{Group: *info.GroupName, Song: *info.SongName}
	song, err := s.repo.GetSongByName(ctx, &songRequest)
	if err != nil {
		log.Printf("ошибка создания песни: %v\n", err)
		return nil, err
	}
	return song, nil
}
func (s *service) DeleteSongByName(ctx context.Context, info *model.SongRequest) (bool, error) {
	err := s.repo.DeleteSongByName(ctx, info)
	if err != nil {
		if err.Error() == "song not found" {
			log.Printf("песня с group %s и song_name %s не найдена", info.Group, info.Song)
			return false, nil
		}
		log.Printf("ошибка удаления песни group %s и song_name %s: %v\n", info.Group, info.Song, err)
		return false, err
	}
	return true, nil
}

func (s *service) GetAllSongs(ctx context.Context, limit int, offset int) ([]*model.SongDetailResponse, error) {
	songs, err := s.repo.GetAllSongs(ctx, limit, offset)
	if err != nil {
		log.Printf("ошибка получения всех песен: %v\n", err)
		return nil, err
	}
	return songs, nil
}

func (s *service) GetSongByName(ctx context.Context, info *model.SongRequest) (*model.SongDetailResponse, error) {
	song, err := s.repo.GetSongByName(ctx, info)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("песня с group %s и song_name %s не найдена", info.Group, info.Song)
			return nil, nil
		}
		log.Printf("ошибка получения песни с group %s и song_name %s: %v\n", info.Group, info.Song, err)
		return nil, err
	}
	return song, nil
}

func (s *service) GetSongText(ctx context.Context, info *model.SongRequest) (*model.SongText, error) {
	song, err := s.repo.GetSongByName(ctx, info)
	if err != nil {
		return nil, err
	}
	if song.Text == "" {
		return nil, nil
	}
	songText := model.SongText{}
	strs := strings.Split(song.Text, "\n")
	cnt := 1
	coupletText := make([]string, 0)
	for _, str := range strs {
		if str != "" {
			coupletText = append(coupletText, str)
		} else {
			couplet := model.Couplet{Couplet: cnt, Text: coupletText}
			songText.Text = append(songText.Text, couplet)
			coupletText = make([]string, 0)
			cnt++
		}
	}
	couplet := model.Couplet{Couplet: cnt, Text: coupletText}
	songText.Text = append(songText.Text, couplet)
	return &songText, nil
}
func (s *service) GetSongsByFilter(ctx context.Context, info *model.SongDetailRequest, limit int, offset int) ([]*model.SongDetailResponse, error) {
	songs, err := s.repo.GetSongsByFilter(ctx, info, limit, offset)
	if err != nil {
		return nil, err
	}
	return songs, nil
}
