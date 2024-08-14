package controllers

import (
	"errors"
	"net/http"
	"server-ssm/db"
	"server-ssm/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPostById(c *gin.Context) {
    postId := c.Param("id")
    var post models.Post

    result := db.DB.Preload("User").First(&post, postId)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            c.JSON(http.StatusNotFound, gin.H{"error": "Posts not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get post"})
        }
        return
    }

    postResponse := post.ToResponse()

    c.JSON(http.StatusOK, gin.H{
        "post": postResponse,
    })
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	result := db.DB.Preload("User").Order("created_at desc").Find(&posts)
	if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get post"})
        }
        return
    }

	postResponses := make([]models.PostResponse, len(posts))
	for i, post := range posts {
		postResponses[i] = post.ToResponse()
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": postResponses,
	})
}

func CreatePost(c *gin.Context) {
	var body struct {
		Content string
	}
	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
		return
	}

	user := c.MustGet("user").(models.User)
	post := models.Post{
		Content: body.Content,
		UserID:  user.ID,
		User:  	 user,
	}
	result := db.DB.Create(&post)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post created",
		"post": post.ToResponse(),
	})
}

func EditPost(c *gin.Context) {
	var body struct {
		Content string
	}
	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
		return
	}

	postId := c.Param("id")
	var post models.Post
	result := db.DB.Preload("User").First(&post, postId)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if post.User.ID != c.MustGet("user").(models.User).ID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	post.Content = body.Content
	result = db.DB.Save(&post)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post updated",
		"post": post.ToResponse(),
	})
}

func DeletePost(c *gin.Context) {
	postId := c.Param("id")
	var post models.Post
	result := db.DB.Preload("User").First(&post, postId)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if post.User.ID != c.MustGet("user").(models.User).ID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	result = db.DB.Delete(&post)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post deleted",
	})
}
