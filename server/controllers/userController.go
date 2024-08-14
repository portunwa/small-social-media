package controllers

import (
	"net/http"
	"server-ssm/db"
	"server-ssm/models"

	"github.com/gin-gonic/gin"
)

func GetUserDetailById(c *gin.Context) {
	userId := c.Param("id")
	var user models.User
	result := db.DB.First(&user, userId)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user.ToResponse(),
	})
}

func GetCurrentUserDetail(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	c.JSON(http.StatusOK, gin.H{
		"user": user.ToResponse(),
	})
}

func EditDisplayName(c *gin.Context) {
	var body struct {
		DisplayName string
	}
	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
		return
	}

	user := c.MustGet("user").(models.User)
	user.DisplayName = body.DisplayName
	result := db.DB.Save(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated",
		"user": user.ToResponse(),
	})
}

func DeleteUser(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	result := db.DB.Delete(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete user"})
		return
	}

	db.DB.Where("user_id = ?", user.ID).Delete(&models.Post{})
	c.SetCookie("token-ssm", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted",
	})
}