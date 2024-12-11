package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shreekumar2901/url-shortener/config"
	"github.com/shreekumar2901/url-shortener/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	DB *gorm.DB
}

var DB DbInstance

func Connect() {
	p := config.Config("POSTGRES_PORT")

	dbPort, err := strconv.Atoi(p)

	if err != nil {
		fmt.Println("Some error occured during DB connect")
	}

	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.Config("POSTGRES_HOST"), config.Config("POSTGRES_USER"),
		config.Config("POSTGRES_PASSWORD"), config.Config("POSTGRES_DB_NAME"), dbPort)

	db, err := gorm.Open(postgres.Open(dbInfo), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database \n", err)
		os.Exit(2)
	}

	log.Println("DB Connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	db.AutoMigrate(&domain.User{})

}
