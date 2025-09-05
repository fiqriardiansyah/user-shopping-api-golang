package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbname     = os.Getenv("DB_NAME")
	dbpassword = os.Getenv("DB_PASSWORD")
	dbusername = os.Getenv("DB_USERNAME")
	dbport     = os.Getenv("DB_PORT")
	dbhost     = os.Getenv("DB_HOST")
	DBInstance *gorm.DB
)

func NewDB() *gorm.DB {
	if DBInstance != nil {
		return DBInstance
	}

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbhost, dbusername, dbpassword, dbname, dbport)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic("❌ database not connected: " + err.Error())
	}

	DBInstance = db

	return DBInstance
}
