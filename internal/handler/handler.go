package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"

	"EffectiveMobile/internal/service"
)

// Handler -
type Handler struct {
	services *service.Service
	logger   *logrus.Logger
}

// NewHandler -
func NewHandler(services *service.Service, logger *logrus.Logger) *Handler {
	return &Handler{
		services: services, logger: logger,
	}
}

// InitRoutes -
func (h *Handler) InitRoutes() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/swagger/", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/swagger/doc.json")))
	router.HandleFunc("GET /info/all", h.getAllSongs)
	router.HandleFunc("POST /info/filter", h.getSongsByFilter)
	router.HandleFunc("GET /info", h.getSongByName)
	router.HandleFunc("GET /info/text", h.getSongTextByName)
	router.HandleFunc("DELETE /info/delete", h.deleteSongByName)
	router.HandleFunc("PATCH /info", h.updateSongByName)
	router.HandleFunc("POST /info", h.createSong)
	return router
}
