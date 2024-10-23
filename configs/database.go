package configs

import (
	"database/sql"
	"fmt"
	"github.com/ssjlee93/fitworks-backend/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbName   = "recordings"
)

func DatabaseConnection() *gorm.DB {

	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	sqlDB, err := sql.Open("pgx", sqlInfo)
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	//db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)

	log.Println("Successfully connected to database")
	return db
}
