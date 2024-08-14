package routes

import (
	"server-ssm/controllers"
	"server-ssm/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.Engine) {
	auth := router.Group("/api/auth")
	{
		auth.POST("/signup", controllers.Signup)
		auth.POST("/login", controllers.Login)
		auth.GET("/logout", middlewares.RequireAuth, controllers.Logout)
		auth.GET("/me", middlewares.RequireAuth, controllers.GetPosts)
	}
}
