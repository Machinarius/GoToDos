package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initialize() {
	connectionString := os.Getenv("POSTGRES_CONNECTION")
	// ⚠️ Be careful of shadowing the global variable!
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}
	migrateStore(db)

	DB = db
}

func migrateStore(db *gorm.DB) {
	db.AutoMigrate(&TodoRecord{})
}
