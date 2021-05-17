package utils

import (
	"promo-code-api/db/promocodes"
	"gorm.io/gorm"
)


func RunMigration(db *gorm.DB){
  		db.AutoMigrate(promocodes.Promocode{}) // database migration

}