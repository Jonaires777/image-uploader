package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// Auth routes
	server.POST("/auth/register", registerUser)
	server.POST("/auth/login", loginUser)
}
