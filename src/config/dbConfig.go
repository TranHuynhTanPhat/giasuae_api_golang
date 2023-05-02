package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Faild to load env file")
	}
	DNS := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOSTNAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"))
	DB, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to database")
	}
	return DB
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close database connection")
	}
	dbSQL.Close()
}
