package config

import (
	"fmt"
	"log"
	"os"

	"github.com/azevedoguigo/thermosync-api/internal/domain"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

func LoadConfig() *DBConfig {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found!")
		log.Panicln("Please, add .env file in root directory.")
	}

	return &DBConfig{
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
	}
}

func InitDB() *gorm.DB {
	cfg := LoadConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err.Error())
	}

	db.AutoMigrate(&domain.User{})

	return db
}
