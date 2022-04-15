package config

import (
	"fmt"
	"log"
	"os"
	"rest/entity"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func GoDotenv(key string) string {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatalf("error loading env file")
	}
	return os.Getenv(key)
}

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error

	dsn := GoDotenv("DB_USERNAME") + ":" +
		GoDotenv("DB_PASSWORD") + "@tcp" +
		"(" + GoDotenv("DB_HOST") + ":" +
		GoDotenv("DB_PORT") + ")/" +
		GoDotenv("DB_NAME") + "?" +
		"parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&entity.Address{}, &entity.User{})
	if err != nil {
		fmt.Printf("Error connecting to database : error=%v", err)
		return nil
	}

	return db
}
