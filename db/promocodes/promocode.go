package promocodes

import (
	"promo-code-api/db"
)

// Promocode struct
type Promocode struct {
	ID                int64
	Name              string  `form:"name" json:"name" example:"barone" format:"string" binding:"required"`
	DateFrom          string  `form:"dateFrom" json:"dateFrom" example:"1/12" format:"string" binding:"required"`
	DateTo            string  `form:"dateTo" json:"dateTo" example:"5/12" format:"string" binding:"required"`
	QuantityAvailable int64   `form:"quantityAvailable" json:"quantityAvailable" example:"5" format:"string" binding:"required"`
	Amount            float64 `form:"amount" json:"amount" example:"5.54" format:"string" binding:"required"`
	QuantityAllocated int64   `form:"quantityAllocated" json:"quantityAllocated" example:"12" format:"string" binding:"required"`
}

// controllers
// first point hit after the route

func GetAllPromoCodes() []Promocode {
	db := db.GetDatabase()
	var records []Promocode
	db.Find(&records)
	return records
}

func GetPromoCodeByID(id int64) Promocode {
	var promocode Promocode
	db := db.GetDatabase()
	db.Where("id = ?", id).Find(&promocode)
	return promocode
}

func checkIfDataExistsInDb(name string) bool {
	var promocode Promocode
	db := db.GetDatabase()
	result := db.Where("name = ?", name).Find(&promocode)
	if result.RowsAffected == 1 {
		return true
	} else {
		return false
	}
}

func AddDataToDb(data Promocode) (string, Promocode) {
	db := db.GetDatabase()
	if !checkIfDataExistsInDb(data.Name) {
		db.Create(&data)
		return "Data successfully added.", data
	} else {
		return "Data already exists, do you want to update it rather?", data
	}

}

func UpdateDataInDbByName(id int64, data Promocode) (string, Promocode) {
	db := db.GetDatabase()
	db.Where("id = ?", id)
	db.Where("id = ?", id).Save(&data)
	return "Data updated", data
}

func DeleteDataFromDb(id int64) string {
	var promocode Promocode
	db := db.GetDatabase()
	db.Where("id = ?", id).Find(&promocode)
	db.Where("id = ?", id).Delete(&promocode)
	return "Data deleted"
}

func BuyPromoCode(id int64) (Promocode,string){
	var promocode Promocode
	db := db.GetDatabase()
	db.Where("id = ?", id).Find(&promocode)
	if(promocode.QuantityAvailable != 0){
		promocode.QuantityAvailable--
		db.Where("id = ?", id).Save(&promocode)
		return promocode, "Promo code purchased."
	}else{
		return promocode, "Promo code no longer available."
	}

}
