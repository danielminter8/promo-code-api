package db

import (
	"fmt"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var db_conn *gorm.DB

func Setup() *gorm.DB{
	dsn := "host=" + os.Getenv("DB_HOST") + " user="+ os.Getenv("DB_USER") + " password="+os.Getenv("DB_PASSWORD") + " dbname="+os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable TimeZone=Africa/Johannesburg"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}else{
		db_conn = db
	}
	return db
}


func GetDatabase() *gorm.DB {
	return db_conn
}
