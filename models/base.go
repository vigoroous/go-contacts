package models

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	// username := os.Getenv("db_user")
	// password := os.Getenv("db_pass")
	dbName := os.Getenv("DB_NAME")
	// dbHost := os.Getenv("db_host")

	//dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	//fmt.Println(dbUri)

	// conn, err := gorm.Open("postgres", dbUri)
	conn, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Account{}, &Contact{})
}

func GetDB() *gorm.DB {
	return db
}
