package main

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", Register)
	r.POST("/login", Login)
}
