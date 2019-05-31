package utils

import (
	"os"
	"fmt"
	"github.com/jinzhu/gorm"
	_"github.com/sirupsen/logrus"

	"github.com/sirupsen/logrus"
)

var db *gorm.DB
var (
	host     = "localhost"
	port     = 5432
	user     = os.Getenv("USER") // your data connection here
	password = os.Getenv("PASS")
	dbname   = os.Getenv("DBNAME")
)
var log = logrus.New()

func InitDB() {
	log.Out = os.Stdout
	file, err := os.OpenFile("session.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file")
	}
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = gorm.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal("dbGorm: failed to connect database: %s", dbinfo)
		db = nil
	} else {
		log.Info("dbGorm: connected with database: %s", dbinfo)
	}

	db.LogMode(true)

}

func GetDB() *gorm.DB {
	if db == nil {
		InitDB()
	}
	return db
}