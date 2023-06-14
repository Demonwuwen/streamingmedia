package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": "success",
	})
}

func Login(c *gin.Context) {
	uname := c.Query("user_name")
	if uname == "" {
		c.Error(errors.New("Query user name error"))
	}
	c.JSON(http.StatusOK, gin.H{
		"name": uname,
	})
}
