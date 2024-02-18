package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/raprast/go-auth/controllers"
	"github.com/raprast/go-auth/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.SetTrustedProxies(nil)
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
