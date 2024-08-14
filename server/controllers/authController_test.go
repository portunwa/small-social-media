package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"server-ssm/controllers"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)


func setupAuthTest() {
	gin.SetMode(gin.TestMode)
}

func TestSignup(t *testing.T) {
	setupAuthTest()

	router := gin.Default()
	router.POST("/signup", controllers.Signup)

	t.Run("failure - invalid body", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/signup", nil)
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestLogin(t *testing.T) {
	setupAuthTest()

	router := gin.Default()
	router.POST("/login", controllers.Login)

	t.Run("failure - invalid body", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/login", nil)
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

}

func TestLogout(t *testing.T) {
	setupAuthTest()

	router := gin.Default()
	router.POST("/logout", controllers.Logout)

	req, _ := http.NewRequest(http.MethodPost, "/logout", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Logout success")
}