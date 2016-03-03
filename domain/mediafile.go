package domain

import (
	"time"
)

type MediaFile struct {
	Id          string
	Path        string
	Title       string
	Album       string
	Artist      string
	AlbumArtist string
	AlbumId     string `parent:"album"`
	TrackNumber int
	DiscNumber  int
	Year        int
	Size        string
	Suffix      string
	Duration    int
	BitRate     int
	Genre       string
	Compilation bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type MediaFileRepository interface {
	BaseRepository
	Put(m *MediaFile) error
	FindByAlbum(albumId string) ([]MediaFile, error)
}
