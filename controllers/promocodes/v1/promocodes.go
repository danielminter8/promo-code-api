package v1

import (
	"fmt"
	"net/http"
	"promo-code-api/db/models"
	"promo-code-api/db/promocodes"

	"github.com/gin-gonic/gin"
)

// GetAllPromoCodes gets all the promocodes
// @Summary Retrieves all promo codes
// @Produce json
// @Success 200 {object} models.Promocode
// @Router /all [get]
func GetAllPromoCodes(r *gin.Context) {
	data := promocodes.GetAllPromoCodes()
	r.JSON(http.StatusOK, gin.H{
		"message": "all promocodes",
		"data":    data,
		"status":  http.StatusOK,
	})
}

// AddPromoCode godoc
// @Summary Add promo code
// @Accept application/json
// @Produce json
// @Param name,quantityAllocated,quantityAvailable,dateFrom,dateTo,amount body models.Promocode true "Example Data"
// @Success 200 {object} models.Promocode
// @Router /add [post]
func AddPromoCode(r *gin.Context) {

	var promocode models.Promocode
	err := r.ShouldBindJSON(&promocode)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid parameters",
		})
	} else {
		message, data := promocodes.AddDataToDb(promocode)
		r.JSON(http.StatusOK, gin.H{
			"message": message,
			"data":    data,
			"status":  http.StatusOK,
		})
	}

}

// UpdatePromoCode godoc
// @Summary Update promo code by name
// @Accept application/json
// @Produce json
// @Param name path string true "name"
// @Param quantityAllocated,quantityAvailable,dateFrom,dateTo,amount body models.Promocode true "Example Data"
// @Success 200 {object} models.Promocode
// @Router /update/{name} [patch]
func UpdatePromoCode(r *gin.Context) {
	var promocode models.Promocode
	name := r.Param("name")
	promocode.Name = name
	err := r.ShouldBindJSON(&promocode)
	fmt.Println(err)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters, make sure you provide dateFrom, dateTo, quanityAllocated, amount, quantityAvailable",
		})
	} else {
		message, data := promocodes.UpdateDataInDbByName(name, promocode)
		r.JSON(http.StatusOK, gin.H{
			"message": message,
			"data":    data,
			"status":  http.StatusOK,
		})
	}

}

// DeletePromoCode godoc
// @Summary Delete promo code by name
// @Produce json
// @Param name path string true "name"
// @Router /delete/{name} [Delete]
func DeletePromoCode(r *gin.Context) {
	name := r.Param("name")
	if name == "" {
		r.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid parameters",
		})
	} else {
		message := promocodes.DeleteDataFromDb(name)
		r.JSON(http.StatusOK, gin.H{
			"message": message,
			"data":    nil,
			"status":  http.StatusOK,
		})
	}

}
