package routes

import (
	"github.com/aishwaryapatle/qastudio/internal/auth/handler"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, authHandler *handler.AuthHandler) {
	authRoutes := rg.Group("/auth")
	{
		authRoutes.POST("/signup", authHandler.Signup)
		authRoutes.POST("/login", authHandler.Login)
	}
}
