package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type User struct {
	ID      uint    `gorm:"primaryKey"`
	Name    string  `gorm:"not null;unique"`
	Age     int     `gorm:"not null"`
	Profile Profile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Profile struct {
	ID                uint `gorm:"primaryKey"`
	UserID            uint `gorm:"not null;unique"`
	Bio               string
	ProfilePictureURL string
}

func connectDB() *gorm.DB {
	dsn := "user=aiiisana password=mypassword dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("failed to configure sql.DB: ", err)
	}

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(0)

	fmt.Println("Successfully connected to the database with GORM!")
	return db
}

func autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&User{}, &Profile{})
	if err != nil {
		log.Fatal("failed to auto-migrate models: ", err)
	}
	fmt.Println("AutoMigrate complete!")
}

func insertUserWithProfile(db *gorm.DB) {
	user := User{
		Name: "John Doe",
		Age:  30,
		Profile: Profile{
			Bio:               "Software developer",
			ProfilePictureURL: "http://example.com/johndoe.jpg",
		},
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatal("failed to insert user with profile: ", err)
	}

	fmt.Println("User and profile inserted successfully!")
}

func queryUserWithProfile(db *gorm.DB) {
	var users []User
	err := db.Preload("Profile").Find(&users).Error
	if err != nil {
		log.Fatal("failed to query users with profiles: ", err)
	}

	fmt.Println("Users and their profiles:")
	for _, user := range users {
		fmt.Printf("User: %s, Age: %d, Profile Bio: %s, Picture: %s\n",
			user.Name, user.Age, user.Profile.Bio, user.Profile.ProfilePictureURL)
	}
}

func updateUserProfile(db *gorm.DB, userID uint, newBio string, newProfilePictureURL string) {
	err := db.Model(&Profile{}).Where("user_id = ?", userID).
		Updates(Profile{Bio: newBio, ProfilePictureURL: newProfilePictureURL}).Error

	if err != nil {
		log.Fatal("failed to update profile: ", err)
	}

	fmt.Println("User profile updated successfully!")
}

func deleteUserWithProfile(db *gorm.DB, userID uint) {
	err := db.Delete(&User{}, userID).Error
	if err != nil {
		log.Fatal("failed to delete user and profile: ", err)
	}

	fmt.Println("User and profile deleted successfully!")
}

func main() {
	db := connectDB()

	autoMigrate(db)

	insertUserWithProfile(db)

	queryUserWithProfile(db)

	updateUserProfile(db, 1, "Updated Bio", "http://example.com/updated.jpg")

	deleteUserWithProfile(db, 1)
}
