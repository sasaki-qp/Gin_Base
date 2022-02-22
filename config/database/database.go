package database

import (
	"log"
	"errors"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = ConnectDB()
	if err != nil {
		log.Print("DEBUG: Database connection Error  ", err)
	}
	return
}

func ConnectDB() (*gorm.DB, error) {
	USER := "golangtest"
	PASS := "golangtest"
	DBNAME := "golangtest"
	PROTOCOL := "tcp(golangtest.cv7rlmmgw3kh.ap-northeast-1.rds.amazonaws.com:3306)"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME

	var db *gorm.DB
	var err error
	db, err = gorm.Open(mysql.Open(CONNECT), &gorm.Config{})

	if err != nil {
		log.Print("DEBUG: Database Connection Error", err)
		return db, errors.New("DB接続に失敗しました")
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB")

	return db, nil
}