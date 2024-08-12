package socket

import (
	"github.com/gin-gonic/gin"
	"server-lu/handler/http"
)

func HttpRouter(g *gin.Engine) {
	router := g.Group("/http")
	{
		router.POST("/health", http.SingletonHealthHandler().HealthHandler)
	}
}
