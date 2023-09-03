package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigEnv struct {
	DB       *gorm.DB
	UserDB   string
	PassDB   string
	HostDB   string
	NameDB   string
	PortDB   string
	PortAPP  string
	PortGRPC string
}

func InitDB() *ConfigEnv {

	godotenv.Load(".env")

	var configEnv ConfigEnv

	configEnv.UserDB = os.Getenv("DB_USER")
	configEnv.PassDB = os.Getenv("DB_PASS")
	configEnv.HostDB = os.Getenv("DB_HOST")
	configEnv.NameDB = os.Getenv("DB_NAME")
	configEnv.PortDB = os.Getenv("DB_PORT")
	configEnv.PortAPP = os.Getenv("PORT")
	configEnv.PortGRPC = os.Getenv("PORT_GRPC")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configEnv.UserDB,
		configEnv.PassDB,
		configEnv.HostDB,
		configEnv.PortDB,
		configEnv.NameDB,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	configEnv.DB = db

	return &configEnv
}
