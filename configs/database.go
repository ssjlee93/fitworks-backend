package configs

import (
	"fmt"
	"github.com/ssjlee93/fitworks-backend/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "ssjlee93"
	password = "postgres"
	dbName   = "recordings"
)

func DatabaseConnection() *gorm.DB {

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=America/New_York", host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helper.ErrorPanic(err)

	log.Println("Successfully connected to database")
	return db
}
