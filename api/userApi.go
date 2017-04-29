package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vcelinServer/db"
	"fmt"
)



func PostUser(c *gin.Context) {

}

func GetUser(c *gin.Context){
	content := gin.H{"Hello": "World" , "Kappa": "GreyFace"}
	c.JSON(200, content)

}

func UpdateUser(c *gin.Context) {

}

func GetUsers(c *gin.Context) {
	var users [] db.User

	context := db.Database()
	context.Find(&users)
	context.Close()

	fmt.Printf("called method users %s",users)
	if (len(users) <= 0) {
		c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No todo found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "data" : users})

}

func DeleteUser(c *gin.Context) {

}



func Logout(c *gin.Context){

}
