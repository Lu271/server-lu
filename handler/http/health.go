package http

import (
	"github.com/gin-gonic/gin"
	"sync"
)

// 测试模块相关

var (
	_healthOnce    sync.Once
	_healthHandler *HealthHandler
)

type HealthHandler struct{}

func SingletonHealthHandler() *HealthHandler {
	_healthOnce.Do(func() {
		_healthHandler = &HealthHandler{}
	})

	return _healthHandler
}

func (h *HealthHandler) HealthHandler(c *gin.Context) {
	println("进入 Health")
}
