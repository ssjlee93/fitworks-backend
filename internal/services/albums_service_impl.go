package services

import (
	"github.com/ssjlee93/fitworks-backend/dtos/responses"
	"github.com/ssjlee93/fitworks-backend/internal/repositories"
)

type AlbumsServiceImpl struct {
	AlbumsRepository repositories.AlbumsRepository
}

func NewAlbumsServiceImpl(a repositories.AlbumsRepository) AlbumsService {
	return &AlbumsServiceImpl{
		AlbumsRepository: a,
	}
}

func (a AlbumsServiceImpl) FindAll() []responses.Album {
	result, _ := a.AlbumsRepository.FindAll()

	var albums []responses.Album
	for _, value := range result {
		album := responses.Album{
			ID:     value.ID,
			Title:  value.Title,
			Artist: value.Artist,
			Price:  value.Price,
		}
		albums = append(albums, album)
	}
	return albums
}
