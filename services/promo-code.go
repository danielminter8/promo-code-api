package services

type Promocode struct{
	Name string
	Date_from string
	Date_to string
	Quantity_available int64
	Amount float64
	Quantity_allocated int64
}

var promocode Promocode
