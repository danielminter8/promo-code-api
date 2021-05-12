package routers

import (
	"net/http"
	"promo-code-api/services"
	"promo-code-api/utils"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Name               string  `json:"name" example:"barone" format:"string"`
	Date_from          string  `json:"date_from" example:"1/12" format:"string"`
	Date_to            string  `json:"date_to" example:"5/12" format:"string"`
	Quantity_available int64   `json:"quantity_available" example:"5" format:"int64"`
	Amount             float64 `json:"amount" example:"5.54" format:"float64"`
	Quantity_allocated int64   `json:"quantity_allocated" example:"12" format:"int64"`
}

// ListAllPromoCodes godoc
// @Summary Retrieves all promo codes
// @Produce json
// @Success 200 {object} Data
// @Router /all [get]
func (c *Controller) GetAllPromoCodes(r *gin.Context) {
	data := services.GetAllPromoCodes()
	r.JSON(http.StatusOK, gin.H{
		"message": "api routes test",
		"data":   data,
		"status":  http.StatusOK,
	})
}

// AddPromoCode godoc
// @Summary Retrieves promo code by name
// @Accept multipart/form-data
// @Produce json
// @Router /add [post]
func (c *Controller) AddPromoCode(r *gin.Context) { // get above swagger 👆 to add body data for example request 
	name := r.PostForm("name")
	date_from := r.PostForm("date_from")
	date_to := r.PostForm("date_to")
	// below strings are converted to int and float
	quantity_available := utils.StringToInt(r.PostForm("quantity_available"))
	amount := utils.StringToFloat(r.PostForm("amount"))
	quantity_allocated := utils.StringToInt(r.PostForm("quantity_allocated"))

	var data services.Promocode
	data.Name = name
	data.Date_from = date_from
	data.Date_to = date_to
	data.Quantity_allocated = quantity_allocated
	data.Quantity_available = quantity_available
	data.Amount = amount
	message, data := services.AddDataToDb(data)

	r.JSON(http.StatusOK, gin.H{
		"message": message,
		"data":    data,
		"status":  http.StatusOK,
	})
}

// UpdatePromoCode godoc
// @Summary Update promo code by name
// @Produce json
// @Param name path string true "name"
// @Router /update/{name} [patch]
func (c *Controller) UpdatePromoCode(r *gin.Context) {
	name := r.Param("name")
	date_from := r.PostForm("date_from")
	date_to := r.PostForm("date_to")
	// below strings are converted to int and float
	quantity_available := utils.StringToInt(r.PostForm("quantity_available"))
	amount := utils.StringToFloat(r.PostForm("amount"))
	quantity_allocated := utils.StringToInt(r.PostForm("quantity_allocated"))

	var data services.Promocode
	data.Name = name
	data.Date_from = date_from
	data.Date_to = date_to
	data.Quantity_allocated = quantity_allocated
	data.Quantity_available = quantity_available
	data.Amount = amount
	message, data := services.UpdateDataInDbByName(name, data)

	r.JSON(http.StatusOK, gin.H{
		"message": message,
		"data":    data,
		"status":  http.StatusOK,
	})
}

// DeletePromoCode godoc
// @Summary Update promo code by name
// @Produce json
// @Param name path string true "name"
// @Router /delete/{name} [post]
func (c *Controller) DeletePromoCode(r *gin.Context) {
	name := r.Param("name")
	message := services.DeleteDataFromDb(name)

	r.JSON(http.StatusOK, gin.H{
		"message": message,
		"data":    nil,
		"status":  http.StatusOK,
	})

}
