package player

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"path"

	"piplayer/model"
)

func (p *Player) GetAlbums() ([]model.Album, error) {
	var albums []model.Album

	artists, err := ioutil.ReadDir(p.dir)
	if err != nil {
		return nil, err
	}
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
			albums = append(albums, model.Album{
				ID:     fmt.Sprintf("%x", md5.Sum([]byte(path.Join(artist.Name(), artistAlbum.Name())))),
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
