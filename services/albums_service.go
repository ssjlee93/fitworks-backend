package services

import "github.com/ssjlee93/fitworks-backend/dtos/responses"

type AlbumsService interface {
	FindAll() []responses.Album
}
