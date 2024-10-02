package routes

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/dert-ops/mediCat/mediCat-Dev.git/cmd/controller"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/:id", controller.GetUserProfile)
	router.PUT("/:id", controller.UpdateUserProfile)
	router.DELETE("/:id", controller.DeleteUserProfile)
	router.PUT("/:id/reset-password", controller.ResetUserPassword)
}
