package models

// Promocode struct
type Promocode struct {
	Name              string  `form:"name" json:"name" example:"barone" format:"string" binding:"required"`
	DateFrom          string  `form:"dateFrom" json:"dateFrom" example:"1/12" format:"string" binding:"required"`
	DateTo            string  `form:"dateTo" json:"dateTo" example:"5/12" format:"string" binding:"required"`
	QuantityAvailable int64   `form:"quantityAvailable" json:"quantityAvailable" example:"5" format:"string" binding:"required"`
	Amount            float64 `form:"amount" json:"amount" example:"5.54" format:"string" binding:"required"`
	QuantityAllocated int64   `form:"quantityAllocated" json:"quantityAllocated" example:"12" format:"string" binding:"required"`
}
