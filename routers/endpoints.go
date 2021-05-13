package routers

import (
	promocodes "promo-code-api/controllers/v1"
	_ "promo-code-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Promo Code API
// @version 1.0
// @description This is a simple api server using the gin web framework.
// @host localhost:5000
// @BasePath /api/v1
func Run() {
	r := gin.Default()

	c := NewController()

	v1Grouping := r.Group("/api/v1")

	{
		v1Grouping.GET("all", promocodes.GetAllPromoCodes)
		v1.POST("add", c.AddPromoCode)
		v1.PATCH("update/:name", c.UpdatePromoCode)
		v1.POST(":id", c.DeletePromoCode)

	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":5000")
}

// v2.GET("all", v1.GetAllPromoCodes())
