package config

import (
	"log"
	"taskmanager/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	dsn := "user=aiiisana password=mypassword dbname=mydb sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	db.Migrator().DropTable(&models.User{})
	db.Migrator().DropTable(&models.Task{})

	err = db.AutoMigrate(&models.Task{}, &models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}
