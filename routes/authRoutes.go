package routes

import (
	"github.com/gin-gonic/gin"
	"joelovien/computer-based-test/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.Signup())
	incomingRoutes.POST("users/signin", controllers.Login())
}
