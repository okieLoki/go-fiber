package database

import (
	"log"
	"os"

	"github.com/okieLoki/go-fiber/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.sqlite"), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database: ", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to database successfully")

	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Migrating database...")

	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	Database = DbInstance{
		Db: db,
	}
}
