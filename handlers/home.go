package handlers

import (
	"net/http"

	"github.com/go-chi/chi"

	"piplayer/model"
	"piplayer/views"
)

type albumsGetter interface {
	GetAlbums() ([]model.Album, error)
}

func Home(mux chi.Router, repo albumsGetter) {
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		albums, err := repo.GetAlbums()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}

		_ = views.Page(views.PageProps{
			Title: "Albums",
			Body:  views.HomeBody(albums),
		}).Render(w)
	})
}
