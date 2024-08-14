package routes

import (
	"server-ssm/controllers"
	"server-ssm/middlewares"

	"github.com/gin-gonic/gin"
)

func PostRoute(router *gin.Engine) {
	post := router.Group("/api/post")
	{
		post.GET("/", controllers.GetPosts)
		post.GET("/:id", middlewares.RequireAuth, controllers.GetPostById)
		post.POST("/create", middlewares.RequireAuth, controllers.CreatePost)
		post.PUT("/:id", middlewares.RequireAuth, controllers.EditPost)
		post.DELETE("/:id", middlewares.RequireAuth, controllers.DeletePost)
	}
}
