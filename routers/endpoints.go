package routers

import (
	"os"
	v1 "promo-code-api/controllers/promocodes/v1"
	_ "promo-code-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Promo Code API
// @version 1.0
// @description This is a simple api server using the gin web framework.
// @host localhost:8080
// @BasePath /api/v1
func Run() {
	r := gin.Default()

	v1Grouping := r.Group("/api/v1")

	{
		v1Grouping.GET("all", v1.GetAllPromoCodes)
		v1Grouping.POST("add", v1.AddPromoCode)
		v1Grouping.PATCH("update/:name", v1.UpdatePromoCode)
		v1Grouping.DELETE("/delete/:name", v1.DeletePromoCode)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("GIN_PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
