package model

import (
	"time"
	"strconv"
	"github.com/TomDev24/MusicServiceGo/internal/database"
	"github.com/TomDev24/MusicServiceGo/pkg/pagination"
)

type Song struct {
	Id			int			`json:"id" gorm:"primaryKey"`
	Group		string		`json:"group"`
	Song		string		`json:"song"`
	ReleaseDate	time.Time	`json:"releaseDate"`
	Text		string		`json:"text"`
	Link		string		`json:"link"`
}

func AllSong(filterFields *Song, db *database.Database, pagination *pagination.Pagination) (*pagination.Pagination, error) {
	var songs []Song

	err := db.Conn.Where(filterFields).Find(&songs).Error
	if err != nil {
		return pagination, err
	}
	db.Conn.Scopes(database.Paginate(songs, pagination, db.Conn)).Find(&songs)
    pagination.Rows = songs

	return pagination, err
}

func GetSongById(db *database.Database, id string, song *Song) error {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
    return db.Conn.First(song, idInt).Error
}

func SaveSong(db *database.Database, song *Song) error {
	return db.Conn.Save(song).Error
}

func DeleteSong(db *database.Database, id string, song *Song) error {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return db.Conn.Delete(song, idInt).Error
}
