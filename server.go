package server

import (
	"net/http"
	"time"
)

// Server -
type Server struct {
	httpServer *http.Server
}

// Run server
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:        ":" + port,
		Handler:     handler,
		ReadTimeout: 5 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}
