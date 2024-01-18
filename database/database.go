package database

import (
	"database/sql"
	"log"
	"mpesa/errors"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *sql.DB
var Db *gorm.DB

func InitDb() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("Initilizing Database")

	//dsn := "root:96E{Q\\iZ/2z$@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:root*@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	errors.HandleErr(err)
	sqlDB, er := db.DB()
	errors.HandleErr(er)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	Db = db
	DB = sqlDB
}
