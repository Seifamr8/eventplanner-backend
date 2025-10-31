package main

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/register", Register)
		api.POST("/login", Login)
	}
}
