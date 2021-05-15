package v1

import (
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
// @Summary Retrieves promo code by name
// @Accept multipart/form-data
// @Produce json
// @Param name,quantityAllocated,quantityAvailable,dateFrom,dateTo,amount body models.Promocode true "Example Data"
// @Success 200 {object} models.Promocode
// @Router /add [post]
func AddPromoCode(r *gin.Context) { // get above swagger 👆 to add body data for example request
	
	var promocode models.Promocode
	err := r.Bind(&promocode)
	// fmt.Println(promocode)
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
// @Produce json
// @Param name path string true "name"
// @Router /update/{name} [patch]
func UpdatePromoCode(r *gin.Context) {
	var promocode models.Promocode
	err := r.Bind(&promocode)

	name := r.Param("name")

	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid parameters",
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
// @Summary Update promo code by name
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
