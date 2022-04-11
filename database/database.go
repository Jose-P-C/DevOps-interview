package database

import (
	"log"
	"os"

	"github.com/Jose-P-C/DevOps-interview/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("courses.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database \n", err.Error())
		os.Exit(2)
	}

	log.Println("Succesfully connected to database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migration")

	// TODO: Add Migration
	db.AutoMigrate(&models.Course{})

	Database = DbInstance{Db: db}
}
