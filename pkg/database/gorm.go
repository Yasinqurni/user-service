package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, string) {

	godotenv.Load(".env")

	dbUser := os.Getenv("DB_USER")
	dbRootPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	Port := os.Getenv("PORT")
	sslmode := os.Getenv("SSL")
	timezone := os.Getenv("TIMEZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbHost, dbUser, dbRootPass, dbName, dbPort, sslmode, timezone)
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbRootPass, dbHost, dbPort, dbName)
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db, Port
}
