package routes

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/dert-ops/mediCat/mediCat-Dev.git/cmd/controller"
	"gitlab.com/dert-ops/mediCat/mediCat-Dev.git/cmd/middleware"
)

func UserRoutes(router *gin.Engine) {
	// Kullanıcı kaydı
	router.POST("/signup", controller.SignUp)

	// Kullanıcı girişi
	router.POST("/signin", controller.SignIn)

	prodected := router.Group("/", middleware.UserServiceAuthMiddleware())
	{

		// Kullanıcı bilgilerini alma
		prodected.GET("/:id", controller.GetUser)

		// Kullanıcı bilgilerini güncelleme
		prodected.PUT("/:id", controller.UpdateUser)

		// Kullanıcıyı silme
		prodected.DELETE("/:id", controller.DeleteUser)

		// Sifre Degistirme
		prodected.PUT("/:id/reset-password", controller.ResetUserPassword)

		// Tüm kullanıcıları listeleme
		prodected.GET("/all", controller.ListUsers)
	}
}
