package v1

import (
	"fmt"
	"net/http"
	"promo-code-api/db/promocodes"
	"promo-code-api/utils"

	"github.com/gin-gonic/gin"
)

// GetAllPromoCodes gets all the promocodes
// @Summary Retrieves all promo codes
// @Produce json
// @Success 200 {object} promocodes.Promocode
// @Router /promo [get]
func GetAllPromoCodes(r *gin.Context) {
	data := promocodes.GetAllPromoCodes()
	r.JSON(http.StatusOK, gin.H{
		"message":   "all promocodes",
		"promocode": data,
	})
}

// GetAllPromoCodes gets all the promocodes
// @Summary Retrieves all promo codes
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} promocodes.Promocode
// @Router /promo/admin/{id} [get]
func GetPromoByID(r *gin.Context) {
	id := r.Param("id")
	convID := utils.StringToInt(id) //convert string id to int64
	data := promocodes.GetPromoCodeByID(convID)
	r.JSON(http.StatusOK, gin.H{
		"message":   "return promocode by id:" + id,
		"promocode": data,
	})

}

// AddPromoCode godoc
// @Summary Add promo code
// @Accept application/json
// @Produce json
// @Param name,quantityAllocated,quantityAvailable,dateFrom,dateTo,amount body promocodes.Promocode true "Example Data"
// @Success 200 {object} promocodes.Promocode
// @Router /promo/admin [post]
func AddPromoCode(r *gin.Context) {
	var promocode promocodes.Promocode
	err := r.ShouldBindJSON(&promocode)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid parameters",
		})
	} else {
		message, data := promocodes.AddDataToDb(promocode)
		r.JSON(http.StatusOK, gin.H{
			"message":   message,
			"promocode": data,
		})
	}

}

// UpdatePromoCode godoc
// @Summary Update promo code by id
// @Accept application/json
// @Produce json
// @Param id path string true "id"
// @Param quantityAllocated,quantityAvailable,dateFrom,dateTo,amount body promocodes.Promocode true "Example Data"
// @Success 200 {object} promocodes.Promocode
// @Router /promo/admin/{id} [patch]
func UpdatePromoCode(r *gin.Context) {
	var promocode promocodes.Promocode
	id := r.Param("id")
	convID := utils.StringToInt(id) //convert string id to int64
	promocode.ID = convID
	err := r.ShouldBindJSON(&promocode)
	fmt.Println(err)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing parameters, make sure you provide dateFrom, dateTo, quanityAllocated, amount, quantityAvailable",
		})
	} else {
		message, data := promocodes.UpdateDataInDbByName(convID, promocode)
		r.JSON(http.StatusOK, gin.H{
			"message":   message,
			"promocode": data,
		})
	}

}

// DeletePromoCode godoc
// @Summary Delete promo code by id
// @Produce json
// @Param id path string true "id"
// @Router /promo/admin/{id} [delete]
func DeletePromoCode(r *gin.Context) {
	id := r.Param("id")
	convID := utils.StringToInt(id) //convert string id to int64
	if id == "" {
		r.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid parameters",
		})
	} else {
		message := promocodes.DeleteDataFromDb(convID)
		r.JSON(http.StatusOK, gin.H{
			"message": message,
			"data":    nil,
		})
	}

}

// BuyPromoCode godoc
// @Summary Should minus from quantity available as it was purchased
// @Produce json
// @Param name path string true "id"
// @Router /promo/app/{id} [post]
func BuyPromoCode(r *gin.Context) {
	id := r.Param("id")

	if id == "" {
		r.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid parameters",
		})
	} else {
		convID := utils.StringToInt(id) //convert string id to int64
		data, message := promocodes.BuyPromoCode(convID)
		r.JSON(http.StatusOK, gin.H{
			"message": message,
			"data":    data,
		})
	}
}
