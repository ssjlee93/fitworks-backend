package services

import (
	"github.com/ssjlee93/fitworks-backend/internal/dtos/responses"
)

type AlbumsService interface {
	FindAll() []responses.Album
}
