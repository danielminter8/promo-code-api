package models


// Promocode struct
type Promocode struct {
	Name              string  `form:"name" json:"name" example:"barone" format:"string"`
	DateFrom          string  `form:"dateFrom" json:"dateFrom" example:"1/12" format:"string"`
	DateTo            string  `form:"dateTo" json:"dateTo" example:"5/12" format:"string"`
	QuantityAvailable int64   `form:"quantityAvailable" json:"quantityAvailable" example:"5" format:"int64"`
	Amount            float64 `form:"amount" json:"amount" example:"5.54" format:"float64"`
	QuantityAllocated int64   `form:"quantityAllocated" json:"quantityAllocated" example:"12" format:"int64"`
}