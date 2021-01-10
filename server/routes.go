package server

import (
	"github.com/go-chi/chi/middleware"

	"piplayer/handlers"
)

func (s *Server) setupRoutes() {
	s.mux.Use(middleware.Recoverer)

	handlers.Static(s.mux)
	handlers.Home(s.mux, s.player)
	handlers.Play(s.mux, s.player)
	handlers.Pause(s.mux, s.player)
}
