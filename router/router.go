package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ssjlee93/fitworks-backend/controllers"
	"net/http"
)

func NewRouter(albumController *controllers.AlbumController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")
	albumRouter := router.Group("/album")
	albumRouter.GET("", albumController.GetAlbums)
	//albumRouter.GET("/:albumId", albumController.FindById)

	return service
}
