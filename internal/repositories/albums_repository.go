package repositories

import (
	"github.com/ssjlee93/fitworks-backend/internal/entities"
)

type AlbumsRepository interface {
	FindAll() ([]entities.Album, error)
	//FindById(id int) (*entities.Album, error)
}
