package promocodes

import "promo-code-api/db"
import "promo-code-api/db/models"


var promocode models.Promocode

// controllers
// first point hit after the route

func GetAllPromoCodes() []models.Promocode {
	db := db.GetDatabase()
	var records []models.Promocode
	db.Find(&records)
	return records
}

func checkIfDataExistsInDb(name string) bool {
	db := db.GetDatabase()
	result := db.Where("name = ?", name).Find(&promocode)
	if result.RowsAffected == 1 {
		return true
	} else {
		return false
	}
}

func AddDataToDb(data models.Promocode) (string, models.Promocode) {
	db := db.GetDatabase()
	if !checkIfDataExistsInDb(data.Name) {
		db.Create(&data)
		return "Data successfully added.", data
	} else {
		return "Data already exists, do you want to update it rather?", data
	}

}

func UpdateDataInDbByName(name string, data models.Promocode) (string, models.Promocode) {
	db := db.GetDatabase()
	db.Where("Name = ?", name)
	// need to add feature to stop program from making data blank if nothing is specified.
	db.Where("Name = ?", name).Save(&data)
	return "Data updated", data
}

func DeleteDataFromDb(name string) string {
	db := db.GetDatabase()
	db.Where("Name = ?", name).Find(&promocode)
	db.Where("Name = ?", name).Delete(&promocode)
	return "Data deleted"
}
