package database

import (
	"finalassignment/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var (
// 	host     = os.Getenv("PGHOST")
// 	port     = os.Getenv("PGPORT")
// 	user     = os.Getenv("PGUSER")
// 	password = os.Getenv("PGPASSWORD")
// 	dbname   = os.Getenv("PGDATABASE")
// )

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
)

var (
	db  *gorm.DB
	err error
)

func GetDB() *gorm.DB {
	return db
}

// @title Final Assignment API
// @version 1.0
// @description This is sample API
// @termsOfService http://google.com/
// @contact.name API Support
// @contact.email soberkode.com
// @license.name Apace 2.0
// @license.url http://google.com/
// @host localhost:8080
// @BasePath /
func StartDB() {
	psqInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err = gorm.Open(postgres.Open(psqInfo), &gorm.Config{})
	fmt.Println(psqInfo)
	if err != nil {
		log.Fatal("Error connecting to database : ", err)
	}

	// Harus urut dari hubungan asosiasi nya T^T
	db.Debug().AutoMigrate(models.Users{}, models.Photos{}, models.SocialMedias{}, models.Comments{})
}
