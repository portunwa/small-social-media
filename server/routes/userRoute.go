package routes

import (
	"server-ssm/controllers"
	"server-ssm/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	user := router.Group("/api/user")
	{
		user.GET("/me", middlewares.RequireAuth, controllers.GetCurrentUserDetail)
		user.GET("/:id", middlewares.RequireAuth, controllers.GetUserDetailById)
		user.PUT("/", middlewares.RequireAuth, controllers.EditDisplayName)
		user.DELETE("/", middlewares.RequireAuth, controllers.DeleteUser)
	}
}
