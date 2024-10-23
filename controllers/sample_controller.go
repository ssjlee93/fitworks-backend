package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ssjlee93/fitworks-backend/dtos/responses"
	"github.com/ssjlee93/fitworks-backend/services"
	"net/http"
)

type AlbumController struct {
	albumService services.AlbumsService
}

func NewAlbumController(service services.AlbumsService) *AlbumController {
	return &AlbumController{albumService: service}
}

func (ac *AlbumController) GetAlbums(ctx *gin.Context) {
	albumResponse := ac.albumService.FindAll()

	webResponse := responses.Response{
		Code:   200,
		Status: "Ok",
		Data:   albumResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
