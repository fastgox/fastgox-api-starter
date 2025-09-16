package utils

import (
	"github.com/gin-gonic/gin"
)

// GetRealIP 获取客户端真实IP
func GetRealIP(c *gin.Context) string {
	// 尝试从各种头部获取真实IP
	clientIP := c.GetHeader("X-Forwarded-For")
	if clientIP == "" {
		clientIP = c.GetHeader("X-Real-Ip")
	}
	if clientIP == "" {
		clientIP = c.GetHeader("X-Forwarded-For")
	}
	if clientIP == "" {
		clientIP = c.ClientIP()
	}
	return clientIP
}
