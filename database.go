package main

import (
	"fmt"
	"go-todo/app/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
}

func (d *Database) ConnectDatabase() error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	d.Db = database
	return nil
}

func (d *Database) Migrate() error {
	if err := d.Db.AutoMigrate(&models.Todo{}); err != nil {
		return err
	}

	log.Println("Successfully migrated models")
	return nil
}

func InitDatabase() *Database {
	var database Database
	godotenv.Load()

	database.ConnectDatabase()
	database.Migrate()
	return &database
}
