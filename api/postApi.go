package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"vcelinServer/db"
)



func CreatePost(c *gin.Context) {
	post := db.Post{Message: c.PostForm("message")};
	context := db.Database()
	defer context.Close()
	context.Save(&post)
	context.Close()
	c.JSON(http.StatusCreated, gin.H{"status" : http.StatusCreated, "message" : "Todo item created successfully!", "resourceId": post.ID})
}

func FetchAllPosts(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK})
}

func FetchSinglePost(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK})
}

func UpdatePost(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "message" : "Todo updated successfully!"})
}

func DeletePost(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "message" : "Todo deleted successfully!"})
}