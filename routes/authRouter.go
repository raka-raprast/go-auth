package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/raprast/go-auth/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.SetTrustedProxies(nil)
	incomingRoutes.POST("users/signup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())
}
