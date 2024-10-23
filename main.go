package main

import (
	"github.com/ssjlee93/fitworks-backend/configs"
	"github.com/ssjlee93/fitworks-backend/controllers"
	"github.com/ssjlee93/fitworks-backend/entities"
	"github.com/ssjlee93/fitworks-backend/helper"
	"github.com/ssjlee93/fitworks-backend/repositories"
	"github.com/ssjlee93/fitworks-backend/router"
	"github.com/ssjlee93/fitworks-backend/services"
	"net/http"
	"time"
)

func main() {
	// db config
	//Database
	db := configs.DatabaseConnection()

	db.Table("album").AutoMigrate(&entities.Album{})

	//Init Repository
	albumRepository := repositories.NewAlbumsRepositoryImpl(db)

	//Init Service
	albumService := services.NewAlbumsServiceImpl(albumRepository)

	//Init controller
	albumController := controllers.NewAlbumController(albumService)

	//Router
	routes := router.NewRouter(albumController)

	server := &http.Server{
		Addr:           ":8888",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)

}
