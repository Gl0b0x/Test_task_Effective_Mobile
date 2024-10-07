package song

import (
	"EffectiveMobile/internal/model"
	"EffectiveMobile/internal/repository/song/conventer"
	modelRepo "EffectiveMobile/internal/repository/song/model"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type postgres struct {
	db *sql.DB
}

// NewSongPostgres -
func NewSongPostgres(db *sql.DB) *postgres {
	return &postgres{db: db}
}

func (repo *postgres) CreateSong(ctx context.Context, info *model.SongDetailRequest) error {
	timeStamp := time.Now()
	q := "INSERT INTO songs ("
	args := make([]interface{}, 0, 7)
	totalArgs := 0
	if info.GroupName != nil {
		q += "group_name, "
		totalArgs++
		args = append(args, *info.GroupName)
	}
	if info.SongName != nil {
		q += "song_name, "
		totalArgs++
		args = append(args, *info.SongName)
	}
	if info.ReleaseDate != nil {
		q += "release_date, "
		totalArgs++
		args = append(args, *info.ReleaseDate)
	}
	if info.Text != nil {
		q += "text, "
		totalArgs++
		args = append(args, *info.Text)
	}
	if info.Link != nil {
		q += "link, "
		totalArgs++
		args = append(args, *info.Link)
	}
	q += "created_at, updated_at) VALUES ("
	for i := 0; i < totalArgs; i++ {
		q += fmt.Sprintf("$%d,", i+1)
	}
	q += fmt.Sprintf("$%d, $%d)", totalArgs+1, totalArgs+2)
	args = append(args, timeStamp, timeStamp)
	_, err := repo.db.ExecContext(ctx, q, args...)
	return err
}

func (repo *postgres) DeleteSongByName(ctx context.Context, info *model.SongRequest) error {
	q := "DELETE FROM songs WHERE group_name = $1 AND song_name = $2"
	res, err := repo.db.ExecContext(ctx, q, info.Group, info.Song)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("song not found")
	}
	return nil
}

func (repo *postgres) GetAllSongs(ctx context.Context, limit int, offset int) ([]*model.SongDetailResponse, error) {
	q := "SELECT * FROM songs LIMIT $1 OFFSET $2"
	rows, err := repo.db.QueryContext(ctx, q, limit, offset)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			log.Println(err)
		}
	}(rows)
	var songs []*model.SongDetailResponse
	for rows.Next() {
		var song modelRepo.SongRepo
		if err = rows.Scan(&song.ID, &song.GroupName, &song.SongName, &song.ReleaseDate, &song.Text, &song.Link, &song.CreatedAt, &song.UpdatedAt); err != nil {
			return nil, err
		}
		songs = append(songs, conventer.ToSongDetailFromSongRepo(&song))
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return songs, nil
}

func (repo *postgres) GetSongByName(ctx context.Context, info *model.SongRequest) (*model.SongDetailResponse, error) {
	q := "SELECT * FROM songs WHERE group_name = $1 AND song_name = $2"
	row := repo.db.QueryRowContext(ctx, q, info.Group, info.Song)
	var song modelRepo.SongRepo
	err := row.Scan(&song.ID, &song.GroupName, &song.SongName, &song.ReleaseDate, &song.Text, &song.Link, &song.CreatedAt, &song.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return conventer.ToSongDetailFromSongRepo(&song), nil
}

func (repo *postgres) UpdateSongByName(ctx context.Context, req *model.SongRequest, info *model.SongDetailRequest) (*model.SongDetailResponse, error) {
	q := "UPDATE songs SET"
	args := make([]interface{}, 0, 5)
	argIndex := 0
	if info.GroupName != nil {
		argIndex++
		q += fmt.Sprintf(" group_name= $%d, ", argIndex)
		args = append(args, *info.GroupName)
	}
	if info.SongName != nil {
		argIndex++
		q += fmt.Sprintf(" song_name= $%d, ", argIndex)
		args = append(args, *info.SongName)
	}
	if info.ReleaseDate != nil {
		argIndex++
		q += fmt.Sprintf(" release_date= $%d, ", argIndex)
		args = append(args, *info.ReleaseDate)
	}
	if info.Text != nil {
		argIndex++
		q += fmt.Sprintf("text = $%d, ", argIndex)
		args = append(args, *info.Text)
	}
	if info.Link != nil {
		argIndex++
		q += fmt.Sprintf("link = $%d, ", argIndex)
		args = append(args, *info.Link)
	}
	q += fmt.Sprintf("updated_at= $%d WHERE group_name= $%d AND song_name = $%d", argIndex, argIndex+1, argIndex+2)
	args = append(args, time.Now(), req.Group, req.Song)
	_, err := repo.db.ExecContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	song, err := repo.GetSongByName(ctx, req)
	if err != nil {
		return nil, err
	}
	return song, nil
}

func (repo *postgres) GetSongsByFilter(ctx context.Context, info *model.SongDetailRequest, limit int, offset int) ([]*model.SongDetailResponse, error) {
	q := "SELECT * FROM songs WHERE"
	args := make([]interface{}, 0, 5)
	argIndex := 0
	if info.GroupName != nil {
		argIndex++
		q += fmt.Sprintf(" group_name = $%d AND", argIndex)
		args = append(args, *info.GroupName)
	}
	if info.SongName != nil {
		argIndex++
		q += fmt.Sprintf(" song_name = $%d AND", argIndex)
		args = append(args, *info.SongName)
	}
	if info.ReleaseDate != nil {
		argIndex++
		q += fmt.Sprintf(" release_date = $%d AND", argIndex)
		args = append(args, *info.ReleaseDate)
	}
	if info.Text != nil {
		argIndex++
		q += fmt.Sprintf(" text = $%d AND", argIndex)
		args = append(args, *info.Text)
	}
	if info.Link != nil {
		argIndex++
		q += fmt.Sprintf(" link = $%d AND", argIndex)
		args = append(args, *info.Link)
	}
	if len(args) == 0 {
		q = q[:len(q)-6]
	} else {
		q = q[:len(q)-4]
	}
	q += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex+1, argIndex+2)
	args = append(args, limit, offset)
	rows, err := repo.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
	}(rows)
	var songs []*model.SongDetailResponse
	for rows.Next() {
		var song modelRepo.SongRepo
		if err = rows.Scan(&song.ID, &song.GroupName, &song.SongName, &song.ReleaseDate, &song.Text, &song.Link, &song.CreatedAt, &song.UpdatedAt); err != nil {
			return nil, err
		}
		songs = append(songs, conventer.ToSongDetailFromSongRepo(&song))
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return songs, nil
}
