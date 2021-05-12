package utils

import (
	"fmt"
	"strconv"
)



func StringToInt(data string) int64{
	converted_data, err := strconv.ParseInt(data, 0, 64) // any other number
	if err != nil{ fmt.Println(err) }
	return converted_data
}

func StringToFloat(data string) float64{
	converted_data, err := strconv.ParseFloat(data, 64) // any other number
	if err != nil{ fmt.Println(err) }
	return converted_data
}