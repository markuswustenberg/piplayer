package player

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"path"
	"sort"
	"strings"

	"piplayer/model"
)

func (p *Player) GetAlbums() ([]model.Album, error) {
	var albums []model.Album

	artists, err := ioutil.ReadDir(p.dir)
	if err != nil {
		return nil, err
	}

	// Make sure the artists are sorted without considering case and "The " in the name
	sort.Slice(artists, func(i, j int) bool {
		artist1 := strings.TrimPrefix(artists[i].Name(), "The ")
		artist2 := strings.TrimPrefix(artists[j].Name(), "The ")
		return strings.ToLower(artist1) < strings.ToLower(artist2)
	})

	for _, artist := range artists {
		if !artist.IsDir() {
			continue
		}
		artistAlbums, err := ioutil.ReadDir(path.Join(p.dir, artist.Name()))
		if err != nil {
			return nil, err
		}
		for _, artistAlbum := range artistAlbums {
			if !artistAlbum.IsDir() {
				continue
			}

			id := fmt.Sprintf("%x", md5.Sum([]byte(path.Join(artist.Name(), artistAlbum.Name()))))
			idFromFile, err := ioutil.ReadFile(path.Join(p.dir, artist.Name(), artistAlbum.Name(), "id.txt"))
			if err == nil {
				id = strings.TrimSpace(string(idFromFile))
			}

			albums = append(albums, model.Album{
				ID:     id,
				Name:   artistAlbum.Name(),
				Artist: artist.Name(),
				Path:   path.Join(artist.Name(), artistAlbum.Name()),
			})
		}
	}
	return albums, nil
}

func (p *Player) GetAlbum(id string) (*model.Album, error) {
	albums, err := p.GetAlbums()
	if err != nil {
		return nil, err
	}
	for _, album := range albums {
		if album.ID == id {
			return &album, nil
		}
	}
	return nil, nil
}
