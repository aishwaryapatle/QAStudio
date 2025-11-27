package routes

import (
	"github.com/aishwaryapatle/qastudio/internal/health"
	"github.com/gin-gonic/gin"
)


func RegisterHealthRoutes(rg *gin.RouterGroup) {
	healthRoutes := rg.Group("/health")
	{
		healthRoutes.GET("/", health.NewHandler().Health)
	}
}
