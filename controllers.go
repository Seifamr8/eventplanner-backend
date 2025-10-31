package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Create(&input)
	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully!",
	})
}

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Login endpoint hit!",
	})
}
