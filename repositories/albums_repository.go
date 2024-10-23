package repositories

import (
	"github.com/ssjlee93/fitworks-backend/entities"
)

type AlbumsRepository interface {
	FindAll() ([]entities.Album, error)
	//FindById(id int) (*entities.Album, error)
}
