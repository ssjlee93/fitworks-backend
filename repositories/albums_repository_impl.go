package repositories

import (
	"github.com/ssjlee93/fitworks-backend/entities"
	"github.com/ssjlee93/fitworks-backend/helper"
	"gorm.io/gorm"
	"log"
)

type AlbumsRepositoryImpl struct {
	DB *gorm.DB
}

func NewAlbumsRepositoryImpl(db *gorm.DB) *AlbumsRepositoryImpl {
	return &AlbumsRepositoryImpl{DB: db}
}

func (a AlbumsRepositoryImpl) FindAll() ([]entities.Album, error) {
	var album []entities.Album
	results := a.DB.Table("album").Find(&album)
	helper.ErrorPanic(results.Error)
	log.Println("Successfully fetched all albums %v", results.RowsAffected)
	return album, nil
}
