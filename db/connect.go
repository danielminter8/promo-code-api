package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)




var g_db *gorm.DB;

func ConnectToDatabase(){
	dsn := "host=promo_code_db user=postgres password=pg123 dbname=postgres port=5432 sslmode=disable TimeZone=Africa/Johannesburg"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	fmt.Println(err)
	g_db = db
}

func GetDatabase() *gorm.DB{
	return g_db
}