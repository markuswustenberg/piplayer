package handlers

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"

	"piplayer/model"
)

type player interface {
	GetAlbum(id string) (*model.Album, error)
	Play(ctx context.Context, path string) error
	Pause(ctx context.Context) error
}

func Play(mux chi.Router, player player) {
	mux.Get("/play/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		album, err := player.GetAlbum(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}

		if album == nil {
			http.Error(w, "album not found", http.StatusNotFound)
			return
		}

		if err := player.Play(r.Context(), album.Path); err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	})
}

func Pause(mux chi.Router, player player) {
	mux.Get("/pause", func(w http.ResponseWriter, r *http.Request) {
		if err := player.Pause(r.Context()); err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	})
}
