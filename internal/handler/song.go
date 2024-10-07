package handler

import (
	"EffectiveMobile/internal/model"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

// getAllSongs godoc
// @Summary Получить информацию о всех песнях
// @Description Возвращает список песен с пагинацией (limit и offset).
// @Tags Songs
// @Accept json
// @Produce json
// @Param limit query int false "Максимальное количество возвращаемых песен" default(10) minimum(1) maximum(50)
// @Param offset query int false "Количетство пропускаемых песен" default(0) minimum(0) maximum(25)
// @Success 200 {array} model.SongDetailResponse "Список песен"
// @Failure 500 {string} Internal server error "Internal server error"
// @Failure 404 {string} No songs found "No songs found"
// @Router /info/all [get]
func (h *Handler) getAllSongs(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 0 || limit > 50 {
		limit = 10
	}
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil || offset < 0 || offset > 25 {
		offset = 0
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	songs, err := h.services.GetAllSongs(ctx, limit, offset)
	if err != nil {
		h.logger.Errorf("error get songs from db: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if songs == nil {
		h.logger.Info("no songs")
		http.Error(w, "no songs", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(songs); err != nil {
		h.logger.Errorf("failed to encode data: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// getSongTextByName получает текст песни по названию и имени группы
// @Summary Получить текст песни
// @Description Возвращает текст песни на основе названия песни и группы
// @Tags Songs
// @Accept json
// @Produce json
// @Param group query string true "Имя группы"
// @Param song query string true "Название песни"
// @Success 200 {object} string "Текст песни"
// @Failure 400 {string} string "Некорректный необходимо передать group, song как query param"
// @Failure 404 {string} string "Текст песни не добавлен"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /info/text [get]
func (h *Handler) getSongTextByName(w http.ResponseWriter, r *http.Request) {
	songName := r.URL.Query().Get("song")
	groupName := r.URL.Query().Get("group")
	if songName == "" || groupName == "" {
		h.logger.Infof("bad request get song=%s, group=%s", songName, groupName)
		http.Error(w, "Bad request need song, group query param", http.StatusBadRequest)
		return
	}
	req := model.SongRequest{Group: groupName, Song: songName}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	text, err := h.services.GetSongText(ctx, &req)
	if err != nil {
		h.logger.Errorf("error get song from db: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if text == nil {
		h.logger.Info("no song text")
		http.Error(w, "no song text", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(text); err != nil {
		h.logger.Errorf("failed to encode data: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// getSongByName godoc
// @Summary Получить информацию о песне по группе и названии песни
// @Description Возвращает информацию о песне
// @Tags Songs
// @Produce json
// @Param group query string true "Название группы"
// @Param song query string true "Название песни"
// @Success 200 {object} model.SongDetailResponse "Информация о песне"
// @Failure 400 {string} string "Неверный запрос, отсутствуют обязательные параметры"
// @Failure 404 {string} string "Песня не найдена"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /info [get]
func (h *Handler) getSongByName(w http.ResponseWriter, r *http.Request) {
	songName := r.URL.Query().Get("song")
	groupName := r.URL.Query().Get("group")
	if songName == "" || groupName == "" {
		h.logger.Infof("bad request get song=%s, group=%s", songName, groupName)
		http.Error(w, "Bad request need song, group query param", http.StatusBadRequest)
		return
	}
	req := model.SongRequest{Group: groupName, Song: songName}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	song, err := h.services.GetSongByName(ctx, &req)
	if err != nil {
		h.logger.Errorf("error get song from db: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if song == nil {
		h.logger.Info("song not found")
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(song); err != nil {
		h.logger.Errorf("failed to encode data: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// deleteSongByName godoc
// @Summary Удаляет запись о песне
// @Description Удаляет запись о песне по группе и названию песни.
// @Tags Songs
// @Param group query string true "Название группы для удаления"
// @Param song query string true "Название песни для удаления"
// @Success 200 {string} string "Песня успешна удалена"
// @Failure 400 {string} string "Некорректный запрос, отсуствуют необходимые данные"
// @Failure 404 {string} string "Не найдено - такой песни нет в базе данных"
// @Failure 500 {string} string "Internal server error"
// @Router /info/delete [delete]
func (h *Handler) deleteSongByName(w http.ResponseWriter, r *http.Request) {
	songName := r.URL.Query().Get("song")
	groupName := r.URL.Query().Get("group")
	if songName == "" || groupName == "" {
		h.logger.Infof("bad request delete song=%s, group=%s", songName, groupName)
		http.Error(w, "Bad request need song, group query param", http.StatusBadRequest)
		return
	}
	req := model.SongRequest{Song: songName, Group: groupName}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ok, err := h.services.DeleteSongByName(ctx, &req)
	if err != nil {
		h.logger.Errorf("error delete song from db: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if !ok {
		h.logger.Info("song not found")
		http.Error(w, "Not found", http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("ok"))
	if err != nil {
		h.logger.Errorf("failed to write response: %v", err)
	}
}

// updateSongByName godoc
// @Summary Обновляет данные песни
// @Description Обновлеят данные песни по группе и названию песни. Обновлеят только переданные данные.
// @Tags Songs
// @Accept json
// @Produce json
// @Param song query string true "Название песни для обновления данных"
// @Param group query string true "Название группы для обновления данных"
// @Param newInfo body model.SongDetailRequest true "Новые данные. Будут обновлены только переданные параметры!"
// @Success 200 {object} model.SongDetailResponse "Обновленные данные песни"
// @Failure 400 {string} string "Неверный запрос, отсутсвуют необходимые данные о песне (group, song)"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /info [patch]
func (h *Handler) updateSongByName(w http.ResponseWriter, r *http.Request) {
	songName := r.URL.Query().Get("song")
	groupName := r.URL.Query().Get("group")
	if songName == "" || groupName == "" {
		h.logger.Infof("bad request get song=%s, group=%s", songName, groupName)
		http.Error(w, "Bad request need song, group query param", http.StatusBadRequest)
		return
	}
	req := model.SongRequest{Group: groupName, Song: songName}
	var newInfo model.SongDetailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Errorf("Failed to decode request body: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			h.logger.Errorf("failed to close body: %v", err)
		}
	}(r.Body)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	song, err := h.services.UpdateSongByName(ctx, &req, &newInfo)
	if err != nil {
		h.logger.Errorf("error update song from db: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(song); err != nil {
		h.logger.Errorf("failed to encode data: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// createSong godoc
// @Summary Создание новой песни
// @Description Добавление песни с приведенными параметрами.
// @Tags Songs
// @Accept json
// @Produce json
// @Param songDetail body model.SongDetailRequest true "Информация о песне для создания"
// @Success 200 {object} model.SongDetailResponse "Информация о созданной песни"
// @Failure 400 {string} string "Некорректный запрос - неправильные входные данные"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /info [post]
func (h *Handler) createSong(w http.ResponseWriter, r *http.Request) {
	var req model.SongDetailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Errorf("Failed to decode request body: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if req.SongName == nil || req.GroupName == nil {
		h.logger.Info("invalid request")
		http.Error(w, "Required song and group field in json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	song, err := h.services.CreateSong(ctx, &req)
	if err != nil {
		h.logger.Errorf("error create song from db: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(song); err != nil {
		h.logger.Errorf("failed to encode data: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// getSongsByFilter godoc
// @Summary Получить песни по фильтру
// @Description Возвращает список песен с пагинацией (limit и offset).
// @Tags Songs
// @Accept json
// @Produce json
// @Param limit query int false "Максимальное количество возвращаемых песен" default(10) minimum(1) maximum(50)
// @Param offset query int false "Количетство пропускаемых песен" default(0) minimum(0) maximum(25)
// @Param songDetail body model.SongDetailRequest true "Информация о песни, по которой будет происходить фильтрация"
// @Success 200 {array} model.SongDetailResponse "Список песен по фильтру"
// @Failure 400 {string} string "Некорректный запрос - неправильные входные данные"
// @Failure 404 {string} string "Не найдено - нет песен удовлетворяющих фильтру"
// @Failure 500 {string} string "Внутренняя ошибка сервиса"
// @Router /info/filter [post]
func (h *Handler) getSongsByFilter(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 0 || limit > 50 {
		limit = 10
	}
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil || offset < 0 || offset > 25 {
		offset = 0
	}
	var req model.SongDetailRequest
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Errorf("Failed to decode request body: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	songs, err := h.services.GetSongsByFilter(ctx, &req, limit, offset)
	if err != nil {
		h.logger.Errorf("error get songs by filter from db: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if songs == nil {
		h.logger.Info("songs not found")
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(songs); err != nil {
		h.logger.Errorf("failed to encode data: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
