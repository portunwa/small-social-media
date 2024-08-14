package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"server-ssm/controllers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

)

func setup() *gin.Engine {
	gin.SetMode(gin.TestMode)

	router := gin.Default()

	router.GET("/posts/:id", controllers.GetPostById)
	router.GET("/posts", controllers.GetPosts)
	router.POST("/posts/create", controllers.CreatePost)
	router.PUT("/posts/:id", controllers.EditPost)
	router.DELETE("/posts/:id", controllers.DeletePost)

	return router
}

func TestCreatePost(t *testing.T) {
	router := setup()

	t.Run("failure - invalid body", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/posts/create", nil)
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), "Invalid body")
	})
}

func TestEditPost(t *testing.T) {
	router := setup()

	t.Run("failure - invalid body", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPut, "/posts/1", nil)
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), "Invalid body")
	})

}
