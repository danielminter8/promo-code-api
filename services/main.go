package services

import (
	// "fmt"
	"promo-code-api/db"
)

func GetAllPromoCodes() []Promocode{
	db := db.GetDatabase()
	var records []Promocode
	 db.Find(&records)
	return records
}

func checkIfDataExistsInDb(name string) bool{
	db := db.GetDatabase()
	result := db.Where("Name = ?", name).Find(&promocode)
	if result.RowsAffected == 1 {
		return true
	}else{
		return false
	}
}

func AddDataToDb(data Promocode) (string,Promocode){
	db := db.GetDatabase()
	 if(!checkIfDataExistsInDb(data.Name)){
		db.Create(&data)
	return "Data successfully added.", data
}else{
	return "Data already exists, do you want to update it rather?", data
}
	
}

func UpdateDataInDbByName(name string, data Promocode) (string,Promocode){
	db := db.GetDatabase()
	db.Where("Name = ?", name)
	// need to add feature to stop program from making data blank if nothing is specified.
	db.Where("Name = ?", name).Save(&data)
	return "Data updated", data
}

func DeleteDataFromDb(name string) string{
	db := db.GetDatabase()
	db.Where("Name = ?", name).Find(&promocode)
	db.Where("Name = ?", name).Delete(&promocode)
	return "Data deleted"
}