package routers

import (
	"os"
	v1 "promo-code-api/controllers/promocodes/v1"
	"promo-code-api/db"
	_ "promo-code-api/docs"
	"promo-code-api/utils"

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

	mainDB := db.Setup()
	utils.RunMigration(mainDB)

	v1Grouping := r.Group("/api/v1")
	promocodeRoutes := v1Grouping.Group("promo")

	adminRoutes := promocodeRoutes.Group("admin")
	adminRoutes.GET("", v1.GetAllPromoCodes)
	adminRoutes.GET("/:id", v1.GetPromoByID)
	adminRoutes.POST("", v1.AddPromoCode)
	adminRoutes.PATCH("/:id", v1.UpdatePromoCode)
	adminRoutes.DELETE("/:id", v1.DeletePromoCode)

	appRoutes := promocodeRoutes.Group("app")
	appRoutes.POST("/:id", v1.BuyPromoCode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("GIN_PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
