package socket

import (
	"github.com/gin-gonic/gin"
	"server-lu/handler/http"
)

func HttpRouter(g *gin.Engine) {
	h_router := g.Group("/http")
	{
		h_router.POST("/health", http.SingletonHealthHandler().HealthHandler)
	}
}
