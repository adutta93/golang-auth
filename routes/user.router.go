package routes

import (
	controller "go-auth/controller"
	middleware "go-auth/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticated())
	incomingRoutes.GET("/get-all-user", controller.GetAllUsers())
	incomingRoutes.GET("/get-user/:id", controller.GetUserById())
}
